package service

import (
	"context"
	"log/slog"
	pb "main/genproto/health_analytics"
	"main/storage"
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

func (s *MedicalService) AddMedicalRecord(ctx context.Context, req *pb.AddMedicalRecordRequest) (*pb.AddMedicalRecordResponse, error){
	return nil, nil
}

func (s *MedicalService) GetMedicalRecord(ctx context.Context, req *pb.GetMedicalRecordRequest) (*pb.GetMedicalRecordResponse, error){
	return nil, nil
}

func (s *MedicalService) UpdateMedicalRecord(ctx context.Context, req *pb.UpdateMedicalRecordRequest) (*pb.UpdateMedicalRecordResponse, error){
	return nil, nil
}

func (s *MedicalService) DeleteMedicalRecord(ctx context.Context, req *pb.DeleteMedicalRecordRequest) (*pb.DeleteMedicalRecordResponse, error){
	return nil, nil
}

func (s *MedicalService) ListMedicalRecord(ctx context.Context, req *pb.ListMedicalRecordsRequest) (*pb.ListMedicalRecordsResponse, error){
	return nil, nil
}


