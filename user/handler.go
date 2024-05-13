package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {

	appGroup := app.Group("/api/user")

	appGroup.Post("/", h.CreateUser)
	appGroup.Get("/:id", h.GetUser)
	appGroup.Delete("/:id", h.DeleteUser)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var createUserDto struct {
		Nickname string `json:"nickname"`
		FullName string `json:"fullName"`
	}

	if err := c.BodyParser(&createUserDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	user, err := h.Service.CreateUser(c.Context(), createUserDto.Nickname, createUserDto.FullName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.Service.GetUserByID(c.Context(), objectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.Service.DeleteUser(c.Context(), objectID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
