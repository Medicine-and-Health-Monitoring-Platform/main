package mongosh

import (
	"encoding/json"
	"fmt"
	pb "main/genproto/health_analytics"
	"main/storage"
	"time"
	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MonitoringRepo struct {
	Coll *mongo.Collection
	Red  *redis.Client
}

func NewMonitoring(db *mongo.Database, rdb *redis.Client) storage.IMonitoringStorage {
	return &MonitoringRepo{
		Coll: db.Collection("monitoring"),
		Red:  rdb,
	}
}
func (r *MonitoringRepo) CreateHealthMonitor(ctx context.Context, req *pb.CreateHealthMonitorReq) error {
	value, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal CreateHealthMonitorReq: %w", err)
	}
	err = r.Red.Set(ctx, req.GetUserId(), value, 0).Err()
	if err != nil {
	 	return fmt.Errorf("failed to store data in Redis: %w", err)
	}

	return nil
}

func (r *MonitoringRepo) GetHealthMonitor(ctx context.Context, userID *pb.UserId) (*pb.GetHealthMonitorsRes, error) {
	val, err := r.Red.Get(ctx,userID.GetUserId()).Result()
	fmt.Println(userID.UserId)
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no data found for user_id: %s", userID.GetUserId())
		}
		return nil, fmt.Errorf("failed to get data from Redis: %w", err)
	}

	var res pb.GetHealthMonitorsRes
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &res, nil
}

func (r *MonitoringRepo) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsRequest) (*pb.GenerateHealthRecommendationsResponse, error) {

	recommendations := []*pb.HealthRecommendation{
		{
			Id:                 uuid.New().String(),
			UserId:             req.GetUserId(),
			RecommendationType: "Exercise",
			Description:        "Increase daily walking to 10,000 steps",
			Priority:           2,
		},
		{
			Id:                 uuid.New().String(),
			UserId:             req.GetUserId(),
			RecommendationType: "Diet",
			Description:        "Reduce sugar intake",
			Priority:           1,
		},
	}

	return &pb.GenerateHealthRecommendationsResponse{
		Recommendations: recommendations,
	}, nil
}

func (r *MonitoringRepo) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringRequest) (*pb.GetRealtimeHealthMonitoringResponse, error) {


	filter := bson.M{"user_id": req.GetUserId()}
	opts := options.FindOne().SetSort(bson.M{"recorded_timestamp": -1})

	var latestData pb.WearableData
	err := r.Coll.FindOne(ctx, filter, opts).Decode(&latestData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get realtime health data: %v", err)
	}

	return &pb.GetRealtimeHealthMonitoringResponse{
		UserId:    latestData.UserId,
		DataType:  latestData.DataType,
		DataValue: latestData.DataValue,
		Timestamp: latestData.RecordedTimestamp,
	}, nil
}

func (r *MonitoringRepo) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryRequest) (*pb.GetDailyHealthSummaryResponse, error) {

	startOfDay, _ := time.Parse("2006-01-02", req.GetDate())
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.M{
		"user_id": req.GetUserId(),
		"recorded_timestamp": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	cursor, err := r.Coll.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get daily health data: %v", err)
	}
	defer cursor.Close(ctx)

	summaryData := make(map[string][]byte)

	return &pb.GetDailyHealthSummaryResponse{
		UserId:      req.GetUserId(),
		Date:        req.GetDate(),
		SummaryData: summaryData,
	}, nil
}

func (r *MonitoringRepo) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryRequest) (*pb.GetWeeklyHealthSummaryResponse, error) {

	startDate, _ := time.Parse("2006-01-02", req.GetStartDate())
	endDate, _ := time.Parse("2006-01-02", req.GetEndDate())
	endDate = endDate.Add(24 * time.Hour) // Include the end date

	filter := bson.M{
		"user_id": req.GetUserId(),
		"recorded_timestamp": bson.M{
			"$gte": startDate,
			"$lt":  endDate,
		},
	}

	cursor, err := r.Coll.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get weekly health data: %v", err)
	}
	defer cursor.Close(ctx)

	summaryData := make(map[string][]byte)

	return &pb.GetWeeklyHealthSummaryResponse{
		UserId:      req.GetUserId(),
		StartDate:   req.GetStartDate(),
		EndDate:     req.GetEndDate(),
		SummaryData: summaryData,
	}, nil
}
