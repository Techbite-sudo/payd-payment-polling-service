package main

import (
	"log"
	"net/http"

	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/api"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := api.SetupRouter(cfg)

	log.Printf("Starting Gateway Service on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
