package config

import (
	"os"
)

type Config struct {
	LogLevel       string
	WhatsappConfig WhatsappConfig
	Database       DatabaseConfig
}

type WhatsappConfig struct {
	AccessToken   string
	PhoneNumberID string
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() Config {
	return Config{
		LogLevel: getEnvWithDefault("LOG_LEVEL", "info"),
		WhatsappConfig: WhatsappConfig{
			AccessToken:   os.Getenv("WHATSAPP_ACCESS_TOKEN"),
			PhoneNumberID: os.Getenv("WHATSAPP_PHONE_NUMBER_ID"),
		},
		Database: DatabaseConfig{
			Driver:   getEnvWithDefault("DB_DRIVER", "mysql"),
			Host:     getEnvWithDefault("DB_HOST", "127.0.0.1"),
			Port:     getEnvWithDefault("DB_PORT", "3307"),
			User:     getEnvWithDefault("DB_USER", "root"),
			Password: getEnvWithDefault("DB_PASSWORD", "root"),
			Name:     getEnvWithDefault("DB_NAME", "auto-messenger"),
			SSLMode:  getEnvWithDefault("DB_SSLMODE", "disable"),
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
