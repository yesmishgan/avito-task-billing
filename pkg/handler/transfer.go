package handler

import (
	"cashbox"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) transferMoney(c *fiber.Ctx) error{

	return nil
}

func (h *Handler) writeOff(c *fiber.Ctx) error{
	var input cashbox.Bill
	if err := c.BodyParser(&input); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := h.services.Write(input); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusOK)
	return nil
}
