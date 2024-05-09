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

const envPath = ".env"

func New() Config {
	var cfg Config

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg.Server.Port = os.Getenv("SERVER_PORT")

	cfg.MongoDB.ConnectionURL = os.Getenv("MONGO_CONNECTION_URL")

	return cfg
}
