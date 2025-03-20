package config

import (
	"os"
)

type Config struct {
	LogLevel        string
	CallMeBotConfig CallMeBotConfig
}

type CallMeBotConfig struct {
	PhoneNumber string
	ApiKey      string
}

func Load() Config {
	return Config{
		LogLevel: getEnvWithDefault("LOG_LEVEL", "info"),
		CallMeBotConfig: CallMeBotConfig{
			PhoneNumber: os.Getenv("CMB_NUMBER"),
			ApiKey:      os.Getenv("CMB_API_KEY"),
		},
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
