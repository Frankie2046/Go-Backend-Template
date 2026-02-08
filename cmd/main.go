package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project/internal/api"
	"project/internal/api/middleware"
	"project/internal/config"
	"project/pkg/logger"
)

func main() {
	cfg := config.Load()
	app := fiber.New()

	l, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}
	defer func() { _ = l.Sync() }()

	app.Use(middleware.Logger(l))
	api.RegisterRoutes(app, l)

	l.Info("server start", zap.String("port", cfg.Port))
	log.Fatal(app.Listen(":" + cfg.Port))
}
