package service

import (
	"context"
	"log/slog"
	pb "main/genproto/health_analytics"
	"main/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MedicalService struct {
	pb.UnimplementedHealthAnalyticsServiceServer
	logger *slog.Logger
	repo   storage.IStorage
}

func NewMedicalService(logger *slog.Logger, repo storage.IStorage) *MedicalService {
	return &MedicalService{
		logger: logger,
		repo:   repo,
	}
}

func (s *MedicalService) AddMedicalRecord(ctx context.Context, req *pb.AddMedicalRecordRequest) (*pb.AddMedicalRecordResponse, error) {
	s.logger.Info("AddMedicalRecord called", "user_id", req.GetRecord().GetUserId())

	resp, err := s.repo.MedicalRecords().AddMedicalRecord(ctx, req)
	if err != nil {
		s.logger.Error("Failed to add medical record", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to add medical record: %v", err)
	}

	s.logger.Info("Medical record added successfully", "record_id", resp.GetRecordId())
	return resp, nil
}

func (s *MedicalService) GetMedicalRecord(ctx context.Context, req *pb.GetMedicalRecordRequest) (*pb.GetMedicalRecordResponse, error) {
	s.logger.Info("GetMedicalRecord called", "record_id", req.GetRecordId())

	resp, err := s.repo.MedicalRecords().GetMedicalRecord(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get medical record", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to get medical record: %v", err)
	}

	s.logger.Info("Medical record retrieved successfully", "record_id", req.GetRecordId())
	return resp, nil
}

func (s *MedicalService) UpdateMedicalRecord(ctx context.Context, req *pb.UpdateMedicalRecordRequest) (*pb.UpdateMedicalRecordResponse, error) {
	s.logger.Info("UpdateMedicalRecord called", "record_id", req.GetRecord().GetId())

	resp, err := s.repo.MedicalRecords().UpdateMedicalRecord(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update medical record", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to update medical record: %v", err)
	}

	s.logger.Info("Medical record updated successfully", "record_id", req.GetRecord().GetId())
	return resp, nil
}

func (s *MedicalService) DeleteMedicalRecord(ctx context.Context, req *pb.DeleteMedicalRecordRequest) (*pb.DeleteMedicalRecordResponse, error) {
	s.logger.Info("DeleteMedicalRecord called", "record_id", req.GetRecordId())

	resp, err := s.repo.MedicalRecords().DeleteMedicalRecord(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete medical record", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete medical record: %v", err)
	}

	s.logger.Info("Medical record deleted successfully", "record_id", req.GetRecordId())
	return resp, nil
}

func (s *MedicalService) ListMedicalRecord(ctx context.Context, req *pb.ListMedicalRecordsRequest) (*pb.ListMedicalRecordsResponse, error) {
	s.logger.Info("ListMedicalRecord called", "user_id", req.GetUserId())

	resp, err := s.repo.MedicalRecords().ListMedicalRecord(ctx, req)
	if err != nil {
		s.logger.Error("Failed to list medical records", "error", err)
		return nil, status.Errorf(codes.Internal, "Failed to list medical records: %v", err)
	}

	s.logger.Info("Medical records listed successfully", "count", len(resp.GetRecords()))
	return resp, nil
}
