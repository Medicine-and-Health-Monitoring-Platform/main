package mongosh

import (
	"main/storage"
	pb "main/genproto/health_analytics"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type MonitoringRepo struct {
	Coll *mongo.Collection
}

func NewMonitoring(db *mongo.Database) storage.IMonitoringStorage {
	return &MonitoringRepo{
		Coll: db.Collection("monitoring"),
	}
}

func (r *MonitoringRepo) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsRequest) (*pb.GenerateHealthRecommendationsResponse, error) {
	// Bu funksiya foydalanuvchi ma'lumotlarini tahlil qilib, tavsiyalar yaratishi kerak
	// Misol uchun:
	
	// 1. Foydalanuvchi ma'lumotlarini olish (medical records, lifestyle data, wearable data)
	// 2. Ma'lumotlarni tahlil qilish
	// 3. Tavsiyalar yaratish

	// Bu yerda sodda misol keltirilgan:
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
	// Bu funksiya real vaqtda ma'lumotlarni olishi kerak
	// Misol uchun, eng so'nggi wearable data'ni olish

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
	// Bu funksiya kunlik sog'liq ma'lumotlarini yig'ishi va tahlil qilishi kerak
	
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
	// Bu yerda olingan ma'lumotlarni tahlil qilish va summaryData'ni to'ldirish kerak

	return &pb.GetDailyHealthSummaryResponse{
		UserId:      req.GetUserId(),
		Date:        req.GetDate(),
		SummaryData: summaryData,
	}, nil
}

func (r *MonitoringRepo) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryRequest) (*pb.GetWeeklyHealthSummaryResponse, error) {
	// Bu funksiya haftalik sog'liq ma'lumotlarini yig'ishi va tahlil qilishi kerak
	
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
	// Bu yerda olingan ma'lumotlarni tahlil qilish va summaryData'ni to'ldirish kerak

	return &pb.GetWeeklyHealthSummaryResponse{
		UserId:      req.GetUserId(),
		StartDate:   req.GetStartDate(),
		EndDate:     req.GetEndDate(),
		SummaryData: summaryData,
	}, nil
}