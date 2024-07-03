package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/errors"
	"github.com/Techbite-sudo/payd-payment-polling-service/common/logger"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/config"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/models"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/payd"
)

type Handler struct {
	Config *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{Config: cfg}
}

func (h *Handler) InitiatePayment(c *gin.Context) {
	var paymentReq models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "Invalid request body", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	paymentResp, err := payd.RunPayment(h.Config.APIUsername, h.Config.APIPassword, paymentReq)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to initiate payment", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusOK, paymentResp)
}

func (h *Handler) GetPaymentStatus(c *gin.Context) {
	transactionID := c.Param("id")

	transactionResp, err := payd.GetTransactionRequests(h.Config.AccountID, h.Config.APIUsername, h.Config.APIPassword)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "Failed to get transaction requests", err)
		logger.Log.Error(appErr)
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	for _, transaction := range transactionResp.TransactionRequests {
		if transaction.ID == transactionID {
			c.JSON(http.StatusOK, transaction)
			return
		}
	}

	appErr := errors.NewAppError(http.StatusNotFound, "Transaction not found", nil)
	logger.Log.Error(appErr)
	c.JSON(appErr.Code, gin.H{"error": appErr.Message})
}
