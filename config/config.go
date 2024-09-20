package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_ROUTER          string
	USER_SERVICE         string
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	SIGNING_KEY          string
	API_GATEWAY          string
	TWIT_SERVICE	     string
	KAFKA_PORT           string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	

	config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ":50050"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgresdb"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "your-password"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "your-db-name"))
	config.USER_ROUTER = cast.ToString(coalesce("USER_ROUTER", ":1596"))
	config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "secret"))
	config.API_GATEWAY = cast.ToString(coalesce("API_GATEWAY", ":4589"))
	config.TWIT_SERVICE = cast.ToString(coalesce("TWIT-SERVICE", ":8082"))
	config.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", "localhost:9092"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
