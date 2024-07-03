package api

import (
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/config"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandler(cfg)

	r.POST("/payments/initiate", h.InitiatePayment)
	r.GET("/payments/status/:id", h.GetPaymentStatus)

	return r
}
