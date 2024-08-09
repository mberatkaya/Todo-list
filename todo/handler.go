package todo

import (
	"TODOproject/utility"
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
	appGroup := app.Group("/api/todo")

	appGroup.Get("/", h.GetAllTodos)
	appGroup.Post("/", h.CreateTodo)
	appGroup.Put("/:id", h.UpdateTodoCompletion)
	appGroup.Delete("/:id", h.DeleteTodo)
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.Service.GetAllTodos(c.Context())
	if err != nil {
		return utility.ErrorResponse(c, err)
	}
	return utility.SuccessResponse(c, todos)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req CreateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c, err)
	}

	todo, err := h.Service.CreateTodo(c.Context(), req.Task)
	if err != nil {
		return utility.ErrorResponse(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(utility.StandardResponse{
		Status: fiber.StatusCreated,
		Data:   todo,
	})
}

func (h *TodoHandler) UpdateTodoCompletion(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c, err)
	}

	var req UpdateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c, err)
	}

	if err := h.Service.UpdateTodoCompletion(c.Context(), objectID, req.Completed); err != nil {
		return utility.ErrorResponse(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c, err)
	}

	if err := h.Service.DeleteTodo(c.Context(), objectID); err != nil {
		return utility.ErrorResponse(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
