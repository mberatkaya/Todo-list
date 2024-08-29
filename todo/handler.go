package todo

import (
	"TODOproject/utility"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	Service TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/api/todo")

	appGroup.Get("/", h.GetAllTodosHandler)
	appGroup.Post("/", h.CreateTodoHandler)
	appGroup.Put("/:id", h.UpdateTodoCompletionHandler)
	appGroup.Delete("/:id", h.DeleteTodoHandler)
}

func (h *TodoHandler) GetAllTodosHandler(c *fiber.Ctx) error {
	todos, err := h.Service.GetAllTodos(c.Context())
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}
	return utility.OkResponse(c, todos)
}

func (h *TodoHandler) CreateTodoHandler(c *fiber.Ctx) error {
	var req CreateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	todo, err := h.Service.CreateTodo(c.Context(), req.Task)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusCreated), todo)
}

func (h *TodoHandler) UpdateTodoCompletionHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	var req UpdateTodoDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	if err := h.Service.UpdateTodoCompletion(c.Context(), objectID, req.Completed); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusOK), nil)
}

func (h *TodoHandler) DeleteTodoHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	if err := h.Service.DeleteTodo(c.Context(), objectID); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusNoContent), nil)
}
