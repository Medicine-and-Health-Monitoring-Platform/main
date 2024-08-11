package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT            string
	AUTH_SERVICE_PORT    string
	PRODUCT_SERVICE_PORT string
	DB_HOST              string
	DB_PORT              int
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	ACCESS_TOKEN         string
	KAFKA_HOST           string
	KAFKA_PORT           string
	KAFKA_TOPIC          string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("/home/jons/go/src/github.com/projects/e-commerce/api-gateway/.env")
		if err != nil {
			log.Fatalf("error loading .env: %v", err)
		}
	}
	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":50051"))
	cfg.PRODUCT_SERVICE_PORT = cast.ToString(coalesce("PRODUCT_SERVICE_PORT", ":8082"))

	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	cfg.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	cfg.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "password"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "auth_i"))

	cfg.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "access_key"))
	cfg.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
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
