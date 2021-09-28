package handler

import (
	"cashbox"
	"github.com/gofiber/fiber/v2"
	"net/http"
)



func (h *Handler) getBalance(c *fiber.Ctx) error{
	var input cashbox.User
	if err := c.BodyParser(&input); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	result, err := h.services.GetBalance(input.Username)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusOK).JSON(result)
	return nil
}
