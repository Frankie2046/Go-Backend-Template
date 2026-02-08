package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project/internal/api/handler"
	"project/internal/api/middleware"
)

func NewApp(logger *zap.Logger, userHandler *handler.UserHandler) *fiber.App {
	app := fiber.New()

	app.Use(middleware.Logger(logger))
	RegisterRoutes(app, userHandler)

	return app
}
