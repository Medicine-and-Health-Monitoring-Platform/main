package mongosh

import (
	"main/config"
	"main/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Storage struct {
	mongo *mongo.Database
}

func ConnectDB() (storage.IStorage, error) {
	opts := options.Client().ApplyURI(config.Load().MongoURI)

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
	}

	return strg, nil
}

func (s Storage) Clouse() {
	s.mongo.Client().Disconnect(context.Background())
}

func (s Storage) MedicalRecords() storage.IMedicalRecordStorage {
	return NewMedecalRecord(s.mongo)
}

func (s Storage) LifestyleData() storage.ILifestyleDataStorage  {
	return NewLifestyleData(s.mongo)
}

func (s Storage) WearableData() storage.IWearableDataStorage {
	return NewWearableData(s.mongo)
}

func (s Storage) Monitoring() storage.IMonitoringStorage {
	return NewMonitoring(s.mongo)
}


