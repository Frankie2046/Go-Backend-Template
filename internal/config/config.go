package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
	AppName  string `yaml:"app_name"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.AppName == "" {
		cfg.AppName = "go-backend"
	}

	return &cfg, nil
}
