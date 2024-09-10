package utility

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data,omitempty"`
	Error  string          `json:"error,omitempty"`
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	response := Response{
		Status: c.Response().StatusCode(),
		Error:  err.Error(),
	}
	return c.Status(response.Status).JSON(response)
}

func OkResponse(c *fiber.Ctx, data interface{}) error {
	dataBytes, _ := json.Marshal(data)
	response := Response{
		Status: c.Response().StatusCode(),
		Data:   dataBytes,
	}
	return c.Status(response.Status).JSON(response)
}
