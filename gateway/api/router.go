package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/config"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/handlers"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/middleware"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandler(cfg)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/payments/initiate", h.InitiatePayment)
		authorized.GET("/payments/status/:id", h.GetPaymentStatus)
	}

	return r
}
