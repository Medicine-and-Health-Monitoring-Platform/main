package mongosh

import (
	"main/storage"
	pb "main/genproto/health_analytics"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LifestyleRepo struct {
	Coll *mongo.Collection
}

func NewLifestyleData(db *mongo.Database) storage.ILifestyleDataStorage {
	return &LifestyleRepo{
		Coll: db.Collection("lifestyle_data"),
	}
}

func (l *LifestyleRepo) AddLifestyleData(ctx context.Context, req *pb.AddLifestyleDataRequest) (*pb.AddLifestyleDataResponse, error) {
	dataID := uuid.New().String()

	data := req.GetData()
	data.Id = dataID

	doc, err := bson.Marshal(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal lifestyle data: %v", err)
	}

	_, err = l.Coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert lifestyle data: %v", err)
	}

	return &pb.AddLifestyleDataResponse{
		DataId: dataID,
	}, nil
}

func (l *LifestyleRepo) GetLifestyleData(ctx context.Context, req *pb.GetLifestyleDataRequest) (*pb.GetLifestyleDataResponse, error) {
	var data pb.LifestyleData

	filter := bson.M{"id": req.GetDataId()}
	err := l.Coll.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Lifestyle data not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get lifestyle data: %v", err)
	}

	return &pb.GetLifestyleDataResponse{
		Data: &data,
	}, nil
}

func (l *LifestyleRepo) UpdateLifestyleData(ctx context.Context, req *pb.UpdateLifestyleDataRequest) (*pb.UpdateLifestyleDataResponse, error) {
	filter := bson.M{"id": req.GetData().GetId()}
	update := bson.M{"$set": req.GetData()}

	result, err := l.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update lifestyle data: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Lifestyle data not found")
	}

	return &pb.UpdateLifestyleDataResponse{
		Success: true,
	}, nil
}

func (l *LifestyleRepo) DeleteLifestyleData(ctx context.Context, req *pb.DeleteLifestyleDataRequest) (*pb.DeleteLifestyleDataResponse, error) {
	filter := bson.M{"id": req.GetDataId()}

	result, err := l.Coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete lifestyle data: %v", err)
	}

	if result.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Lifestyle data not found")
	}

	return &pb.DeleteLifestyleDataResponse{
		Success: true,
	}, nil
}