package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	APIUsername string
	APIPassword string
	AccountID   string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Port:        getEnv("PORT", "8082"),
		APIUsername: getEnv("API_USERNAME", ""),
		APIPassword: getEnv("API_PASSWORD", ""),
		AccountID:   getEnv("ACCOUNT_ID", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
