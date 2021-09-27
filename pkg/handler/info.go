package handler

import (
	"cashbox"
	"github.com/gofiber/fiber/v2"
	"net/http"
)



func (h *Handler) getBalance(c *fiber.Ctx) error{
	var input cashbox.User
	if err := c.BodyParser(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	result, err := h.services.GetBalance(input.Username, "RUB")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	c.Status(http.StatusOK).JSON(result)
	return nil
}
