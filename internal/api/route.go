package api

import (
	"github.com/gofiber/fiber/v2"

	"project/internal/api/handler"
)

func RegisterRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/users/:id", userHandler.GetUser)
}
