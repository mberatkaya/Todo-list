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
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid request payload"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

	user, err := h.Service.CreateUser(c.Context(), req.Nickname, req.FullName, req.Password)
	if err != nil {
		return utility.ErrorResponse(c, err)
	}

<<<<<<< HEAD
	return c.Status(fiber.StatusCreated).JSON(utility.StandardResponse{
		Status: fiber.StatusCreated,
		Data:   user,
	})
=======
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": user})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid ID"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

	user, err := h.Service.GetUserByID(c.Context(), objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
<<<<<<< HEAD
			return c.Status(fiber.StatusNotFound).JSON(utility.StandardResponse{
				Status: fiber.StatusNotFound,
				Data:   "User not found",
			})
=======
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "error": "User not found"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
		}
		return utility.ErrorResponse(c, err)
	}

<<<<<<< HEAD
	return utility.SuccessResponse(c, user)
=======
	return utility.OkResponse(c, user)
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid ID"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

	var req UpdateUserDto
	if err := c.BodyParser(&req); err != nil {
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid request payload"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

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
<<<<<<< HEAD
			return utility.ErrorResponse(c, err)
=======
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": "Failed to hash password"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
		}
		updateFields = append(updateFields, bson.E{"password", string(hashedPassword)})
	}

	if err := h.Service.UpdateUser(c.Context(), objectID, updateFields); err != nil {
		if err.Error() == "nickname already exists" {
<<<<<<< HEAD
			return c.Status(fiber.StatusConflict).JSON(utility.StandardResponse{
				Status: fiber.StatusConflict,
				Data:   "Nickname already exists",
			})
=======
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "error": "Nickname already exists"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
		}
		return utility.ErrorResponse(c, err)
	}

	return utility.OkResponse(c, nil)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid ID"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

	if err := h.Service.DeleteUser(c.Context(), objectID); err != nil {
		return utility.ErrorResponse(c, err)
	}

	return utility.OkResponse(c, nil)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
<<<<<<< HEAD
		return utility.ErrorResponse(c, err)
=======
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid request payload"})
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	}

	user, err := h.Service.ValidatePassword(c.Context(), req.Nickname, req.Password)
	if err != nil {
<<<<<<< HEAD
		return c.Status(fiber.StatusUnauthorized).JSON(utility.StandardResponse{
			Status: fiber.StatusUnauthorized,
			Data:   "Invalid username or password",
		})
	}

	return utility.SuccessResponse(c, user)
=======
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "error": err.Error()})
	}

	return utility.OkResponse(c, user)
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
}
