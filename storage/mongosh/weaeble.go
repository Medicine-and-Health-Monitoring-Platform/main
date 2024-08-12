package mongosh

import (
	"main/storage"
	pb"main/genproto/health_analytics"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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
	return nil, nil
}

func(r *WearableRepo) GetWearableData(ctx context.Context, req *pb.GetWearableDataRequest) (*pb.GetWearableDataResponse, error) {
	return nil, nil
}

func(r *WearableRepo) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataRequest) (*pb.UpdateWearableDataResponse, error) {
	return nil, nil
}

func (r *WearableRepo)DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataRequest) (*pb.DeleteWearableDataResponse, error) {
	return nil, nil
}
