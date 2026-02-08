package config

type Config struct {
	Port     string
	LogLevel string
	AppName  string
}

func Load() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "info",
		AppName:  "go-backend",
	}
}
