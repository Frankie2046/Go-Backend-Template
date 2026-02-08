package main

import (
	"log"

	"go.uber.org/zap"

	"project/internal/wire"
)

func main() {
	app, l, cfg, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("init app failed: %v", err)
	}
	defer func() { _ = l.Sync() }()

	l.Info("server start", zap.String("port", cfg.Port))
	log.Fatal(app.Listen(":" + cfg.Port))
}
