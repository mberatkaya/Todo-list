package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
}

type ServerConfig struct {
	Port string
}

type MongoDBConfig struct {
	ConnectionURL string
	DatabaseName  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
		MongoDB: MongoDBConfig{
			ConnectionURL: os.Getenv("MONGO_CONNECTION_URL"),
			DatabaseName:  os.Getenv("MONGO_DATABASE_NAME"),
		},
	}
}
