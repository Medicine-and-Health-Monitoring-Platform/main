package service


import (
	"context"
	"log/slog"
	pb "main/genproto/health_analytics"
	"main/storage"
)

type WearableService struct {
	pb.UnimplementedHealthAnalyticsServiceServer
	logger *slog.Logger
	repo   storage.IStorage
}

func NewWearableService(logger *slog.Logger, repo storage.IStorage) *WearableService {
	return &WearableService{
		logger: logger,
		repo:   repo,
	}
}

func (s *WearableService) AddWearableData(ctx context.Context, req *pb.AddWearableDataRequest) (*pb.AddWearableDataResponse, error) {
	return nil, nil
}

func (s *WearableService) GetWearableData(ctx context.Context, req *pb.GetWearableDataRequest) (*pb.GetWearableDataResponse, error) {
	return nil, nil
}

func (s *WearableService) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataRequest) (*pb.UpdateWearableDataResponse, error) {
	return nil, nil
}

func (s *WearableService) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataRequest) (*pb.DeleteWearableDataResponse, error) {
	return nil, nil
}