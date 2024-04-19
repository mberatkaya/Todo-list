package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"strconv"
)

// Todo struct'ı, bir To-Do öğesini temsil eder
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type CreateTodo struct {
	Task string `json:"task"`
}

// In-memory veritabanı olarak kullanılacak slice
var todos []Todo
var idCounter = 1

func getAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	var createTodoDto CreateTodo
	if err := c.BodyParser(&createTodoDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var newTodo Todo
	newTodo.ID = idCounter
	newTodo.Task = createTodoDto.Task

	idCounter++
	todos = append(todos, newTodo)

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

func updateTodoCompletion(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedTodo Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Belirtilen ID'ye sahip To-Do öğesini bul
	found := false
	for i, todo := range todos {
		if todo.ID == id {
			// To-Do öğesini güncelle
			todos[i].Completed = updatedTodo.Completed
			found = true
			break
		}
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func deleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(fiber.StatusNoContent).Send(nil)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
}

func main() {
	// Fiber uygulamasını oluştur
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	// Router'lar
	app.Get("/todos", getAllTodos)
	app.Post("/todos", createTodo)
	app.Delete("/todos/:id", deleteTodo)
	app.Put("/todos/:id", updateTodoCompletion)

	// Uygulamayı belirtilen portta başlat
	app.Listen(":8080")
}
