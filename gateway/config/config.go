package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port               string
    AuthServiceURL     string
    PaymentsServiceURL string
}

func Load() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        return nil, err
    }

    return &Config{
        Port:               getEnv("PORT", "8080"),
        AuthServiceURL:     getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
        PaymentsServiceURL: getEnv("PAYMENTS_SERVICE_URL", "http://localhost:8082"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}