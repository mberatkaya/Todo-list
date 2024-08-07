package utility

import (
	"github.com/gofiber/fiber/v2"
)

func OkResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}
