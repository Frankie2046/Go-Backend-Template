package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"project/internal/config"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	level := zapcore.InfoLevel
	if err := level.Set(strings.ToLower(cfg.LogLevel)); err != nil {
		level = zapcore.InfoLevel
	}

	zcfg := zap.NewProductionConfig()
	zcfg.Level = zap.NewAtomicLevelAt(level)
	zcfg.InitialFields = map[string]interface{}{
		"app": cfg.AppName,
	}
	return zcfg.Build()
}
