package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Logger(l *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		l.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("latency", time.Since(start)),
		)
		return err
	}
}
