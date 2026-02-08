//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"project/internal/api"
	"project/internal/api/handler"
	"project/internal/config"
	"project/internal/repo"
	"project/internal/service"
	"project/pkg/logger"
)

func InitializeApp() (*fiber.App, *zap.Logger, *config.Config, error) {
	wire.Build(
		config.Load,
		logger.NewLogger,
		repo.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,
		api.NewApp,
	)
	return nil, nil, nil, nil
}
