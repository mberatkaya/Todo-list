package utility

import "github.com/gofiber/fiber/v2"

type StandardResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(StandardResponse{
		Status: fiber.StatusOK,
		Data:   data,
	})
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(StandardResponse{
		Status: fiber.StatusInternalServerError,
		Data:   err.Error(),
	})
}
