package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) getBalance(c *fiber.Ctx) error{
	c.SendStatus(200)
	return nil
}
