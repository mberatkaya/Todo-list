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
	"TODOproject/user"
)

func main() {
	cfg := config.NewConfig()

	// MongoDB bağlantı ayarları
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.ConnectionURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantı hatası: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Veritabanı ve koleksiyonlar
	db := client.Database("mydatabase") // Database name direkt olarak connection URL'den alınabilir
	todoCollection := db.Collection("todos")
	userCollection := db.Collection("users")

	// Fiber uygulaması
	app := fiber.New()
	app.Use(logger.New())

	// Todo servisi ve handler
	todoRepo := todo.NewTodoRepository(todoCollection)
	todoService := todo.NewTodoService(todoRepo)
	todoHandler := todo.NewTodoHandler(todoService)
	todoHandler.RegisterRoutes(app)

	// User servisi ve handler
	userRepo := user.NewUserRepository(userCollection)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)
	userHandler.RegisterRoutes(app)

	// Sunucu başlatma
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
