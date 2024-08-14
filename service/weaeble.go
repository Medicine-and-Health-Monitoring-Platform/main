package service

import (
	"context"
	pb "main/genproto/health_analytics"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type WearableService struct {
// 	pb.UnimplementedHealthAnalyticsServiceServer
// 	logger *slog.Logger
// 	repo   storage.IStorage
// }

// func NewWearableService(logger *slog.Logger, repo storage.IStorage) *WearableService {
// 	return &WearableService{
// 		logger: logger,
// 		repo:   repo,
// 	}
// }

func (s *HealthService) AddWearableData(ctx context.Context, req *pb.AddWearableDataRequest) (*pb.AddWearableDataResponse, error) {
	s.logger.Info("AddWearableData called", "user_id", req.GetData().GetUserId())

	resp, err := s.repo.WearableData().AddWearableData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to add wearable data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to add wearable data: %v", err)
	}

	s.logger.Info("Wearable data added successfully", "data_id", resp.GetDataId())
	return resp, nil
}

func (s *HealthService) GetWearableData(ctx context.Context, req *pb.GetWearableDataRequest) (*pb.GetWearableDataResponse, error) {
	s.logger.Info("GetWearableData called", "data_id", req.GetDataId())

	resp, err := s.repo.WearableData().GetWearableData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get wearable data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get wearable data: %v", err)
	}

	s.logger.Info("Wearable data retrieved successfully", "data_id", req.GetDataId())
	return resp, nil
}

func (s *HealthService) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataRequest) (*pb.UpdateWearableDataResponse, error) {
	s.logger.Info("UpdateWearableData called", "data_id", req.GetData().GetId())

	resp, err := s.repo.WearableData().UpdateWearableData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update wearable data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to update wearable data: %v", err)
	}

	s.logger.Info("Wearable data updated successfully", "data_id", req.GetData().GetId())
	return resp, nil
}

func (s *HealthService) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataRequest) (*pb.DeleteWearableDataResponse, error) {
	s.logger.Info("DeleteWearableData called", "data_id", req.GetDataId())

	resp, err := s.repo.WearableData().DeleteWearableData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete wearable data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete wearable data: %v", err)
	}

	s.logger.Info("Wearable data deleted successfully", "data_id", req.GetDataId())
	return resp, nil
}
