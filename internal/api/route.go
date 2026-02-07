package api

import (
	"github.com/gofiber/fiber/v2"

	"project/internal/api/handler"
	"project/internal/repo"
	"project/internal/service"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app.Get("/users/:id", userHandler.GetUser)
}