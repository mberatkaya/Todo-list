package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct'ı uygulamanın yapılandırma bilgilerini tutar
type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
}

// ServerConfig struct'ı HTTP sunucu yapılandırma bilgilerini tutar
type ServerConfig struct {
	Port string
}

// MongoDBConfig struct'ı MongoDB yapılandırma bilgilerini tutar
type MongoDBConfig struct {
	ConnectionURL string
	DatabaseName  string
}

// NewConfig fonksiyonu yapılandırma bilgisini döner
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
