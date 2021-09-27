package handler

import (
	"cashbox/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *fiber.App{
	router := fiber.New()
	router.Use(logger.New(logger.Config{
		Format:     "[${time}][${ip}]:$t{port} ${status} - ${method} ${path}\n",
		Output: os.Stdout,
	}))

	api := router.Group("/api")
	{
		info := api.Group("/info")
		{
			info.Post("/balance", h.getBalance)

		}
		transaction := api.Group("/transaction")
		{
			transaction.Post("/transfer", h.transferMoney)
			transaction.Post("/write", h.writeOff)
		}
	}

	return router
}