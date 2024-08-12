package service

import (
	"context"
	"log/slog"
	pb "main/genproto/health_analytics"
	"main/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LifestyleService struct {
	pb.UnimplementedHealthAnalyticsServiceServer
	logger *slog.Logger
	repo   storage.IStorage
}

func NewLifestyleService(logger *slog.Logger, repo storage.IStorage) *LifestyleService {
	return &LifestyleService{
		logger: logger,
		repo:   repo,
	}
}

func (s *LifestyleService) AddLifestyleData(ctx context.Context, req *pb.AddLifestyleDataRequest) (*pb.AddLifestyleDataResponse, error) {
	s.logger.Info("AddLifestyleData called", "user_id", req.GetData().GetUserId())

	resp, err := s.repo.LifestyleData().AddLifestyleData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to add lifestyle data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to add lifestyle data: %v", err)
	}

	s.logger.Info("Lifestyle data added successfully", "data_id", resp.GetDataId())
	return resp, nil
}

func (s *LifestyleService) GetLifestyleData(ctx context.Context, req *pb.GetLifestyleDataRequest) (*pb.GetLifestyleDataResponse, error) {
	s.logger.Info("GetLifestyleData called", "data_id", req.GetDataId())

	resp, err := s.repo.LifestyleData().GetLifestyleData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get lifestyle data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get lifestyle data: %v", err)
	}

	s.logger.Info("Lifestyle data retrieved successfully", "data_id", req.GetDataId())
	return resp, nil
}

func (s *LifestyleService) UpdateLifestyleData(ctx context.Context, req *pb.UpdateLifestyleDataRequest) (*pb.UpdateLifestyleDataResponse, error) {
	s.logger.Info("UpdateLifestyleData called", "data_id", req.GetData().GetId())

	resp, err := s.repo.LifestyleData().UpdateLifestyleData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update lifestyle data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to update lifestyle data: %v", err)
	}

	s.logger.Info("Lifestyle data updated successfully", "data_id", req.GetData().GetId())
	return resp, nil
}

func (s *LifestyleService) DeleteLifestyleData(ctx context.Context, req *pb.DeleteLifestyleDataRequest) (*pb.DeleteLifestyleDataResponse, error) {
	s.logger.Info("DeleteLifestyleData called", "data_id", req.GetDataId())

	resp, err := s.repo.LifestyleData().DeleteLifestyleData(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete lifestyle data", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete lifestyle data: %v", err)
	}

	s.logger.Info("Lifestyle data deleted successfully", "data_id", req.GetDataId())
	return resp, nil
}