package service

import (
	"context"
	pb "main/genproto/health_analytics"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type MonitoringService struct {
// 	pb.UnimplementedHealthAnalyticsServiceServer
// 	logger *slog.Logger
// 	repo   storage.IStorage
// }

// func NewMonitoringService(logger *slog.Logger, repo storage.IStorage) *MonitoringService {
// 	return &MonitoringService{
// 		logger: logger,
// 		repo:   repo,
// 	}
// }

func (s *HealthService) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsRequest) (*pb.GenerateHealthRecommendationsResponse, error) {
	s.logger.Info("GenerateHealthRecommendations called", "user_id", req.GetUserId())

	resp, err := s.repo.Monitoring().GenerateHealthRecommendations(ctx, req)
	if err != nil {
		s.logger.Error("Failed to generate health recommendations", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to generate health recommendations: %v", err)
	}

	s.logger.Info("Health recommendations generated successfully", "recommendations_count", len(resp.GetRecommendations()))
	return resp, nil
}

func (s *HealthService) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringRequest) (*pb.GetRealtimeHealthMonitoringResponse, error) {
	s.logger.Info("GetRealtimeHealthMonitoring called", "user_id", req.GetUserId())

	resp, err := s.repo.Monitoring().GetRealtimeHealthMonitoring(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get realtime health monitoring data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get realtime health monitoring data: %v", err)
	}

	s.logger.Info("Realtime health monitoring data retrieved successfully", "user_id", resp.GetUserId())
	return resp, nil
}

func (s *HealthService) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryRequest) (*pb.GetDailyHealthSummaryResponse, error) {
	s.logger.Info("GetDailyHealthSummary called", "user_id", req.GetUserId(), "date", req.GetDate())

	resp, err := s.repo.Monitoring().GetDailyHealthSummary(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get daily health summary", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get daily health summary: %v", err)
	}

	s.logger.Info("Daily health summary retrieved successfully", "user_id", resp.GetUserId(), "date", resp.GetDate())
	return resp, nil
}

func (s *HealthService) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryRequest) (*pb.GetWeeklyHealthSummaryResponse, error) {
	s.logger.Info("GetWeeklyHealthSummary called", "user_id", req.GetUserId(), "start_date", req.GetStartDate(), "end_date", req.GetEndDate())

	resp, err := s.repo.Monitoring().GetWeeklyHealthSummary(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get weekly health summary", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get weekly health summary: %v", err)
	}

	s.logger.Info("Weekly health summary retrieved successfully", "user_id", resp.GetUserId(), "start_date", resp.GetStartDate(), "end_date", resp.GetEndDate())
	return resp, nil
}
