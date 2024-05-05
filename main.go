package main

import (
	"TODOproject/handler"
	"TODOproject/repository"
	"TODOproject/service"
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
	todoRepo := repository.NewTodoRepository(client, "mydatabase", "todos")
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// Fiber oluşturma
	app := fiber.New()

	// Middleware'ler
	app.Use(logger.New())

	// Router'lar
	app.Get("/todos", todoHandler.GetAllTodos)
	app.Post("/todos", todoHandler.CreateTodo)
	app.Put("/todos/:id", todoHandler.UpdateTodoCompletion)
	app.Delete("/todos/:id", todoHandler.DeleteTodo)

	// Uygulama portu
	log.Fatal(app.Listen(":8080"))
}
