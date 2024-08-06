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

func (h *TodoHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/api/todo") // API rotası için grup oluştur

	appGroup.Get("/", h.GetAllTodos)
	appGroup.Post("/", h.CreateTodo)
	appGroup.Put("/:id", h.UpdateTodoCompletion)
	appGroup.Delete("/:id", h.DeleteTodo)
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.Service.GetAllTodos(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Todos retrieved successfully",
		"data":    todos,
	})
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req CreateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	todo, err := h.Service.CreateTodo(c.Context(), req.Task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Todo created successfully",
		"data":    todo,
	})
}

func (h *TodoHandler) UpdateTodoCompletion(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var req UpdateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := h.Service.UpdateTodoCompletion(c.Context(), objectID, req.Completed); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Todo updated successfully",
	})
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.Service.DeleteTodo(c.Context(), objectID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  "success",
		"message": "Todo deleted successfully",
	})
}
