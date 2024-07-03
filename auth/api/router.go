package api

import (
	"github.com/Techbite-sudo/payd-payment-polling-service/auth/handlers"
	"github.com/Techbite-sudo/payd-payment-polling-service/auth/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, jwtSecret string) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandler(db, jwtSecret)

	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)

	// Example of a protected route
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	{
		protected.GET("/protected", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{
				"message": "This is a protected route",
				"user":    username,
			})
		})
	}

	return r
}
