package mongosh

import (
	"main/storage"
	pb"main/genproto/health_analytics"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)


type MedicalRepo struct {
	Coll *mongo.Collection
}


func NewMedecalRecord(db *mongo.Database) storage.IMedicalRecordStorage {
	return &MedicalRepo{
		Coll: db.Collection("medical_records"),
	}
}

func (r *MedicalRepo) AddMedicalRecord(ctx context.Context,  req *pb.AddMedicalRecordRequest)(*pb.AddMedicalRecordResponse, error){
	return nil, nil
}

func (r *MedicalRepo) GetMedicalRecord(ctx context.Context,  req *pb.GetMedicalRecordRequest)(*pb.GetMedicalRecordResponse, error){
	return nil, nil
}

func (r *MedicalRepo) UpdateMedicalRecord(ctx context.Context,  req *pb.UpdateMedicalRecordRequest)(*pb.UpdateMedicalRecordResponse, error){
	return nil, nil
}

func (r *MedicalRepo) DeleteMedicalRecord(ctx context.Context,  req *pb.DeleteMedicalRecordRequest)(*pb.DeleteMedicalRecordResponse, error){
	return nil, nil
}

func (r *MedicalRepo) ListMedicalRecord(ctx context.Context,  req *pb.ListMedicalRecordsRequest)(*pb.ListMedicalRecordsResponse, error){
	return nil, nil
}


