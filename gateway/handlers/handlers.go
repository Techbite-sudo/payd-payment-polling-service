package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/errors"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/logger"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/config"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/utils"
)

type Handler struct {
	Config *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{Config: cfg}
}

func (h *Handler) Register(c *gin.Context) {
	err := utils.ReverseProxy(h.Config.AuthServiceURL + "/auth/register")(c)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to proxy register request", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}
}

func (h *Handler) Login(c *gin.Context) {
	err := utils.ReverseProxy(h.Config.AuthServiceURL + "/auth/login")(c)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to proxy login request", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}
}

func (h *Handler) InitiatePayment(c *gin.Context) {
	err := utils.ReverseProxy(h.Config.PaymentsServiceURL + "/payments/initiate")(c)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to proxy payment initiation request", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}
}

func (h *Handler) GetPaymentStatus(c *gin.Context) {
	err := utils.ReverseProxy(h.Config.PaymentsServiceURL + "/payments/status")(c)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to proxy payment status request", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}
}
