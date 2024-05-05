package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	Service *TodoService
}

func NewTodoHandler(service *TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.Service.GetAllTodos(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(todos)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var createTodoDto struct {
		Task string `json:"task"`
	}

	if err := c.BodyParser(&createTodoDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	todo, err := h.Service.CreateTodo(c.Context(), createTodoDto.Task)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (h *TodoHandler) UpdateTodoCompletion(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedTodo struct {
		Completed bool `json:"completed"`
	}

	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := h.Service.UpdateTodoCompletion(c.Context(), objectID, updatedTodo.Completed); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.Service.DeleteTodo(c.Context(), objectID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/todo")
	appGroup.Get("/todos", h.GetAllTodos)
	appGroup.Post("/todos", h.CreateTodo)
	appGroup.Put("/todos/:id", h.UpdateTodoCompletion)
	appGroup.Delete("/todos/:id", h.DeleteTodo)
}
