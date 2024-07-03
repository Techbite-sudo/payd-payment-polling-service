package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Techbite-sudo/payd-payment-polling-service/auth/models"
	"gorm.io/gorm"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/errors"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/logger"
)

type Handler struct {
	DB        *gorm.DB
	JWTSecret string
}

func NewHandler(db *gorm.DB, jwtSecret string) *Handler {
	return &Handler{DB: db, JWTSecret: jwtSecret}
}

func (h *Handler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "Invalid request body", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to hash password", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to create user", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "Invalid request body", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	var user models.User
	if err := h.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		appErr := errors.NewAppError(http.StatusUnauthorized, "Invalid credentials", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
		appErr := errors.NewAppError(http.StatusUnauthorized, "Invalid credentials", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to generate token", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
