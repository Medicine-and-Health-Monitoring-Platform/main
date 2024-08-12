package mongosh

import (
	"main/storage"
	pb"main/genproto/health_analytics"


	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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
	return nil, nil
}

func (r *MonitoringRepo) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringRequest) (*pb.GetRealtimeHealthMonitoringResponse, error) {
	return nil, nil
}

func (r *MonitoringRepo) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryRequest) (*pb.GetDailyHealthSummaryResponse, error) {
	return nil, nil
}

func (r *MonitoringRepo) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryRequest) (*pb.GetWeeklyHealthSummaryResponse, error) {
	return nil, nil
}


