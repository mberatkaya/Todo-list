package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Uygulama yapılandırma bilgilerini tutar
type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
}

// HTTP sunucu yapılandırma bilgilerini tutar
type ServerConfig struct {
	Port string
}

// MongoDB yapılandırma bilgilerini tutar
type MongoDBConfig struct {
	ConnectionURL string
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
		},
	}
}
