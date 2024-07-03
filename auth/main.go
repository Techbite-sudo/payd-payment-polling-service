package main

import (
    "net/http"

    "github.com/Techbite-sudo/payd-payment-polling-service/auth/api"
    "github.com/Techbite-sudo/payd-payment-polling-service/auth/config"
    "github.com/Techbite-sudo/payd-payment-polling-service/auth/database"
    "github.com/Techbite-sudo/payd-payment-polling-service/common/logger"
)

func main() {
    logger.Init()

    cfg, err := config.Load()
    if err != nil {
        logger.Log.Fatalf("Failed to load configuration: %v", err)
    }

    db, err := database.InitDB(cfg.DatabaseURL)
    if err != nil {
        logger.Log.Fatalf("Failed to initialize database: %v", err)
    }

    router := api.SetupRouter(db, cfg.JWTSecret)

    logger.Log.Infof("Starting Authentication Service on :%s", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
        logger.Log.Fatalf("Failed to start server: %v", err)
    }
}