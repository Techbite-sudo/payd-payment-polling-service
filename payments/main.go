package main

import (
	"net/http"

	"github.com/Techbite-sudo/payd-payment-polling-service/common/logger"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/api"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/config"
)

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatalf("Failed to load configuration: %v", err)
	}

	router := api.SetupRouter(cfg)

	logger.Log.Infof("Starting Payments Service on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
	}
}
