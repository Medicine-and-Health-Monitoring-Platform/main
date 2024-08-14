package mongosh

import (
	pb "main/genproto/health_analytics"
	"main/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *MedicalRepo) AddMedicalRecord(ctx context.Context, req *pb.AddMedicalRecordRequest) (*pb.AddMedicalRecordResponse, error) {
	recordID := uuid.New().String()

	record := req.GetRecord()
	record.Id = recordID

	doc, err := bson.Marshal(record)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal record: %v", err)
	}

	_, err = r.Coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert record: %v", err)
	}

	return &pb.AddMedicalRecordResponse{
		RecordId: recordID,
	}, nil
}

func (r *MedicalRepo) GetMedicalRecord(ctx context.Context, req *pb.GetMedicalRecordRequest) (*pb.GetMedicalRecordResponse, error) {
	var record pb.MedicalRecord

	filter := bson.M{"id": req.GetRecordId()}
	err := r.Coll.FindOne(ctx, filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Medical record not found")
		}
		return nil, err
	}

	return &pb.GetMedicalRecordResponse{
		Record: &record,
	}, nil
}

func (r *MedicalRepo) UpdateMedicalRecord(ctx context.Context, req *pb.UpdateMedicalRecordRequest) (*pb.UpdateMedicalRecordResponse, error) {
	filter := bson.M{"id": req.GetRecord().GetId()}
	update := bson.M{"$set": req.GetRecord()}

	result, err := r.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Medical record not found")
	}

	return &pb.UpdateMedicalRecordResponse{
		Success: true,
	}, nil
}

func (r *MedicalRepo) DeleteMedicalRecord(ctx context.Context, req *pb.DeleteMedicalRecordRequest) (*pb.DeleteMedicalRecordResponse, error) {
	filter := bson.M{"id": req.GetRecordId()}

	result, err := r.Coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Medical record not found")
	}

	return &pb.DeleteMedicalRecordResponse{
		Success: true,
	}, nil
}

func (r *MedicalRepo) ListMedicalRecord(ctx context.Context, req *pb.ListMedicalRecordsRequest) (*pb.ListMedicalRecordsResponse, error) {
	filter := bson.M{"user_id": req.GetUserId()}

	skip := int64((req.GetPage() - 1) * req.GetPageSize())
	limit := int64(req.GetPageSize())

	totalCount, err := r.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	cursor, err := r.Coll.Find(ctx, filter, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var records []*pb.MedicalRecord
	for cursor.Next(ctx) {
		var record pb.MedicalRecord
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListMedicalRecordsResponse{
		Records:    records,
		TotalCount: int32(totalCount),
	}, nil
}
