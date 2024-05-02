package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Todo struct'ı, bir To-Do öğesini temsil eder
type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task"`
	Completed bool               `json:"completed"`
}

type CreateTodo struct {
	Task string `json:"task"`
}

var client *mongo.Client
var todoCollection *mongo.Collection

func initMongoDB() {
	// MongoDB'ye bağlanma URI'si
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB istemcisini oluştur
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// MongoDB bağlantısını test et
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Veritabanı ve koleksiyon referansını al
	todoCollection = client.Database("mydatabase").Collection("todos")
}

func getAllTodos(c *fiber.Ctx) error {
	// MongoDB'den tüm To-Do öğelerini getir
	cursor, err := todoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	var todos []Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		return err
	}

	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	var createTodoDto CreateTodo
	if err := c.BodyParser(&createTodoDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	newTodo := Todo{
		Task:      createTodoDto.Task,
		Completed: false,
	}

	// MongoDB'ye yeni To-Do öğesini ekle
	result, err := todoCollection.InsertOne(context.Background(), newTodo)
	if err != nil {
		return err
	}

	// Insert işlemi sonucunda oluşan ObjectID'yi Todo struct'ına ata
	newTodo.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

func updateTodoCompletion(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedTodo Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// MongoDB'de belirtilen ID'ye sahip To-Do öğesini güncelle
	filter := bson.D{{"_id", objectID}}
	update := bson.D{{"$set", bson.D{{"completed", updatedTodo.Completed}}}}
	_, err = todoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// MongoDB'de belirtilen ID'ye sahip To-Do öğesini sil
	filter := bson.D{{"_id", objectID}}
	result, err := todoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func main() {
	// MongoDB'ye bağlan
	initMongoDB()

	// Fiber uygulamasını oluştur
	app := fiber.New()

	// Middleware'ler
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	// Router'lar
	app.Get("/todos", getAllTodos)
	app.Post("/todos", createTodo)
	app.Delete("/todos/:id", deleteTodo)
	app.Put("/todos/:id", updateTodoCompletion)

	// Uygulamayı belirtilen portta başlat
	log.Fatal(app.Listen(":8080"))
}
