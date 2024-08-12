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

type WearableRepo struct {
	Coll *mongo.Collection
}

func NewWearableData(db *mongo.Database) storage.IWearableDataStorage {
	return &WearableRepo{
		Coll: db.Collection("wearable_data"),
	}
}

func (r *WearableRepo) AddWearableData(ctx context.Context, req *pb.AddWearableDataRequest) (*pb.AddWearableDataResponse, error) {
	dataID := uuid.New().String()

	data := req.GetData()
	data.Id = dataID

	doc, err := bson.Marshal(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal wearable data: %v", err)
	}

	_, err = r.Coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert wearable data: %v", err)
	}

	return &pb.AddWearableDataResponse{
		DataId: dataID,
	}, nil
}

func (r *WearableRepo) GetWearableData(ctx context.Context, req *pb.GetWearableDataRequest) (*pb.GetWearableDataResponse, error) {
	var data pb.WearableData

	filter := bson.M{"id": req.GetDataId()}
	err := r.Coll.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "Wearable data not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get wearable data: %v", err)
	}

	return &pb.GetWearableDataResponse{
		Data: &data,
	}, nil
}

func (r *WearableRepo) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataRequest) (*pb.UpdateWearableDataResponse, error) {
	filter := bson.M{"id": req.GetData().GetId()}
	update := bson.M{"$set": req.GetData()}

	result, err := r.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update wearable data: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Wearable data not found")
	}

	return &pb.UpdateWearableDataResponse{
		Success: true,
	}, nil
}

func (r *WearableRepo) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataRequest) (*pb.DeleteWearableDataResponse, error) {
	filter := bson.M{"id": req.GetDataId()}

	result, err := r.Coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete wearable data: %v", err)
	}

	if result.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Wearable data not found")
	}

	return &pb.DeleteWearableDataResponse{
		Success: true,
	}, nil
}