package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project/internal/api/handler"
	"project/internal/repo"
	"project/internal/service"
)

func RegisterRoutes(app *fiber.App, logger *zap.Logger) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, logger)

	app.Get("/users/:id", userHandler.GetUser)
}
