package user

import (
	"TODOproject/utility"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	appGroup.Get("/:id", h.GetUserByID)
	appGroup.Put("/:id", h.UpdateUser)
	appGroup.Delete("/:id", h.DeleteUser)
	appGroup.Post("/login", h.Login)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	user, err := h.Service.CreateUser(c.Context(), req.Nickname, req.FullName, req.Password)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusCreated), user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	user, err := h.Service.GetUserByID(c.Context(), objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return utility.ErrorResponse(c.Status(fiber.StatusNotFound), err)
		}
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c, user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	var req UpdateUserDto
	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	// Update fields
	updateFields := bson.D{}
	if req.Nickname != nil {
		updateFields = append(updateFields, bson.E{"nickname", *req.Nickname})
	}
	if req.FullName != nil {
		updateFields = append(updateFields, bson.E{"fullName", *req.FullName})
	}
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
		}
		updateFields = append(updateFields, bson.E{"password", string(hashedPassword)})
	}

	if err := h.Service.UpdateUser(c.Context(), objectID, updateFields); err != nil {
		if err.Error() == "nickname already exists" {
			return utility.ErrorResponse(c.Status(fiber.StatusConflict), err)
		}
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusOK), nil)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	if err := h.Service.DeleteUser(c.Context(), objectID); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusInternalServerError), err)
	}

	return utility.OkResponse(c.Status(fiber.StatusNoContent), nil)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusBadRequest), err)
	}

	user, err := h.Service.ValidatePassword(c.Context(), req.Nickname, req.Password)
	if err != nil {
		return utility.ErrorResponse(c.Status(fiber.StatusUnauthorized), err)
	}

	return utility.OkResponse(c, user)
}
