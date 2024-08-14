package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT         string
	MDB_NAME          string
	MongoURI          string
	DB_HOST           string
	DB_PORT           int
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	ACCESS_TOKEN      string
	KAFKA_HOST        string
	KAFKA_PORT        string
	KAFKA_TOPIC       string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("/home/fedy/go/src/github/PersonalizedMedicineAndHealthMonitoringPlatform/main/.env")
		if err != nil {
			log.Fatalf("error loading .env: %v", err)
		}
	}
	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", "health:8083"))


	

	cfg.MDB_NAME = cast.ToString(coalesce("MDB_NAME", "medicine"))
	cfg.MongoURI = cast.ToString(coalesce("MONGO_URI", "mongodb://mongo:27017"))

	cfg.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "access_key"))
	cfg.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "kafka"))
	cfg.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", "9092"))
	cfg.KAFKA_TOPIC = cast.ToString(coalesce("KAFKA_TOPIC", "e-commerce:order"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
