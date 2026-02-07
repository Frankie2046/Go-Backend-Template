package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"project/internal/api"
	"project/internal/config"
	"project/internal/api/middleware"
)

func main() {
	cfg := config.Load()
	app := fiber.New()

	api.RegisterRoutes(app)
    app.Use(middleware.Logger())
	log.Printf("server start on :%s\n", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}