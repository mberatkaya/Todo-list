package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"TODOproject/config"
	"TODOproject/todo"
)

func main() {
	// Yapılandırmayı al
	cfg := config.New()

	// MongoDB bağlantı ayarları
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.ConnectionURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantı hatası: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Todo Repository, Service ve Handler'larını oluştur
	todoRepo := todo.NewTodoRepository(client, "mydatabase", "todos")
	todoService := todo.NewTodoService(todoRepo)
	todoHandler := todo.NewTodoHandler(todoService)

	// Fiber uygulamasını oluştur
	app := fiber.New()

	// Middleware'ler
	app.Use(logger.New())

	// Router'ları kaydet
	todoHandler.RegisterRoutes(app)

	// Uygulamayı belirtilen port üzerinden dinle
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
