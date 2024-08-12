package mongosh

import (
	"main/storage"
	pb"main/genproto/health_analytics"


	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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
	return nil, nil
}

func (l *LifestyleRepo) GetLifestyleData(ctx context.Context, req *pb.GetLifestyleDataRequest) (*pb.GetLifestyleDataResponse, error) {
	return nil, nil
}

func (l *LifestyleRepo) UpdateLifestyleData(ctx context.Context, req *pb.UpdateLifestyleDataRequest) (*pb.UpdateLifestyleDataResponse, error) {
	return nil, nil
}

func (l *LifestyleRepo) DeleteLifestyleData(ctx context.Context, req *pb.DeleteLifestyleDataRequest) (*pb.DeleteLifestyleDataResponse, error) {
	return nil, nil
}


