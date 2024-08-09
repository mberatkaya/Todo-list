package utility

<<<<<<< HEAD
import "github.com/gofiber/fiber/v2"

type StandardResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(StandardResponse{
		Status: fiber.StatusOK,
		Data:   data,
=======
import (
	"github.com/gofiber/fiber/v2"
)

func OkResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	})
}

func ErrorResponse(c *fiber.Ctx, err error) error {
<<<<<<< HEAD
	return c.Status(fiber.StatusInternalServerError).JSON(StandardResponse{
		Status: fiber.StatusInternalServerError,
		Data:   err.Error(),
=======
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
>>>>>>> e7302d5a5cff8e6414bb3f1cf9a648c7fa80799d
	})
}
