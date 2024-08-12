package service

import (
	"context"
	"log/slog"
	pb "main/genproto/health_analytics"
	"main/storage"
)

type LifestyService struct {
	pb.UnimplementedHealthAnalyticsServiceServer
	logger *slog.Logger
	repo   storage.IStorage
}

func NewLifestyleData(logger *slog.Logger, repo storage.IStorage) *LifestyService {
	return &LifestyService{
		logger: logger,
		repo:   repo,
	}
}

func (s *LifestyService) AddLifestyleData(ctx context.Context, req *pb.AddLifestyleDataRequest) (*pb.AddLifestyleDataResponse, error) {
	return nil, nil
}

func (s *LifestyService) GetLifestyleData(ctx context.Context, req *pb.GetLifestyleDataRequest) (*pb.GetLifestyleDataResponse, error) {
	return nil, nil
}

func (s *LifestyService) UpdateLifestyleData(ctx context.Context, req *pb.UpdateLifestyleDataRequest) (*pb.UpdateLifestyleDataResponse, error) {
	return nil, nil
}

func (s *LifestyService) DeleteLifestyleData(ctx context.Context, req *pb.DeleteLifestyleDataRequest) (*pb.DeleteLifestyleDataResponse, error) {
	return nil, nil
}