package handler

import (
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *fiber.Ctx, statusCode int, message string){
	c.JSON(fiber.Map{

	})
}