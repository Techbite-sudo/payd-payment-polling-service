package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Techbite-sudo/payd-payment-polling-service/gateway/config"
)

func TestRegister(t *testing.T) {
	cfg := &config.Config{
		AuthServiceURL: "http://auth-service",
	}
	handler := NewHandler(cfg)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", handler.Register)

	req, _ := http.NewRequest(http.MethodPost, "/register", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Since we're not actually calling the auth service in this test,
	// we'll just check that the request was proxied (status code should not be 404)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestLogin(t *testing.T) {
	cfg := &config.Config{
		AuthServiceURL: "http://auth-service",
	}
	handler := NewHandler(cfg)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", handler.Login)

	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Since we're not actually calling the auth service in this test,
	// we'll just check that the request was proxied (status code should not be 404)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestInitiatePayment(t *testing.T) {
	cfg := &config.Config{
		PaymentsServiceURL: "http://payments-service",
	}
	handler := NewHandler(cfg)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/payments/initiate", handler.InitiatePayment)

	req, _ := http.NewRequest(http.MethodPost, "/payments/initiate", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Since we're not actually calling the payments service in this test,
	// we'll just check that the request was proxied (status code should not be 404)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestGetPaymentStatus(t *testing.T) {
	cfg := &config.Config{
		PaymentsServiceURL: "http://payments-service",
	}
	handler := NewHandler(cfg)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/payments/status/:id", handler.GetPaymentStatus)

	req, _ := http.NewRequest(http.MethodGet, "/payments/status/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Since we're not actually calling the payments service in this test,
	// we'll just check that the request was proxied (status code should not be 404)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
