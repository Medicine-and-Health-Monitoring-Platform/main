package mongosh

import (
	"main/config"
	"main/storage"

	// "main/storage/redis"
	rdb "main/storage/redis"

	"github.com/redis/go-redis/v9"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Storage struct {
	mongo *mongo.Database
	redis *redis.Client
}

func ConnectDB() (storage.IStorage, error) {
	opts := options.Client().ApplyURI(config.Load().MongoURI).SetAuth(options.Credential{
		Username: "root",
		Password: "example",})

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	strg := Storage{
		mongo: client.Database(config.Load().MDB_NAME),
		redis: rdb.ConnectRDB(),
	}

	return strg, nil
}

func (s Storage) Clouse() {
	s.mongo.Client().Disconnect(context.Background())
}

func (s Storage) MedicalRecords() storage.IMedicalRecordStorage {
	return NewMedecalRecord(s.mongo)
}

func (s Storage) LifestyleData() storage.ILifestyleDataStorage {
	return NewLifestyleData(s.mongo)
}

func (s Storage) WearableData() storage.IWearableDataStorage {
	return NewWearableData(s.mongo)
}

func (s Storage) Monitoring() storage.IMonitoringStorage {
	return NewMonitoring(s.mongo, s.redis)
}
