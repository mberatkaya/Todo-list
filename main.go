package main

import (
	"TODOproject/todo"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB bağlamtısı
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Repository, Service ve Handler'ları oluşturma
	todoRepo := todo.NewTodoRepository(client, "mydatabase", "todos")
	todoService := todo.NewTodoService(todoRepo)
	todoHandler := todo.NewTodoHandler(todoService)

	// Fiber oluşturma
	app := fiber.New()

	// Middleware'ler
	app.Use(logger.New())

	// Router'lar
	todoHandler.RegisterRoutes(app)

	// Uygulama portu
	log.Fatal(app.Listen(":8080"))
}
