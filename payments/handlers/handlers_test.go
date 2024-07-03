package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/config"
	"github.com/Techbite-sudo/payd-payment-polling-service/payments/models"
)

// MockPaydClient is a mock for the Payd API client
type MockPaydClient struct {
	mock.Mock
}

func (m *MockPaydClient) RunPayment(username, password string, paymentReq models.PaymentRequest) (*models.PaymentResponse, error) {
	args := m.Called(username, password, paymentReq)
	return args.Get(0).(*models.PaymentResponse), args.Error(1)
}

func (m *MockPaydClient) GetTransactionRequests(accountID, username, password string) (*models.TransactionResponse, error) {
	args := m.Called(accountID, username, password)
	return args.Get(0).(*models.TransactionResponse), args.Error(1)
}

func TestInitiatePayment(t *testing.T) {
	mockClient := new(MockPaydClient)
	cfg := &config.Config{
		APIUsername: "testuser",
		APIPassword: "testpass",
	}
	handler := NewHandler(cfg)
	handler.paydClient = mockClient

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/payments/initiate", handler.InitiatePayment)

	testCases := []struct {
		name           string
		requestBody    models.PaymentRequest
		mockResponse   *models.PaymentResponse
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful payment initiation",
			requestBody: models.PaymentRequest{
				Username:    "user123",
				NetworkCode: "123",
				Amount:      100.0,
				PhoneNumber: "1234567890",
				Narration:   "Test payment",
				Currency:    "USD",
				CallbackURL: "http://example.com/callback",
			},
			mockResponse: &models.PaymentResponse{
				MerchantRequestID:   "123456",
				CheckoutRequestID:   "789012",
				ResponseCode:        "0",
				ResponseDescription: "Success",
				CustomerMessage:     "Payment initiated",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "Success",
		},
		{
			name:        "Invalid request body",
			requestBody: models.PaymentRequest{
				// Missing required fields
			},
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body",
		},
		{
			name: "Payd API error",
			requestBody: models.PaymentRequest{
				Username:    "user123",
				NetworkCode: "123",
				Amount:      100.0,
				PhoneNumber: "1234567890",
				Narration:   "Test payment",
				Currency:    "USD",
				CallbackURL: "http://example.com/callback",
			},
			mockResponse:   nil,
			mockError:      errors.New("Payd API error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to initiate payment",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient.On("RunPayment", cfg.APIUsername, cfg.APIPassword, tc.requestBody).Return(tc.mockResponse, tc.mockError).Once()

			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/payments/initiate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestGetPaymentStatus(t *testing.T) {
	mockClient := new(MockPaydClient)
	cfg := &config.Config{
		APIUsername: "testuser",
		APIPassword: "testpass",
		AccountID:   "acc123",
	}
	handler := NewHandler(cfg)
	handler.paydClient = mockClient

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/payments/status/:id", handler.GetPaymentStatus)

	testCases := []struct {
		name           string
		transactionID  string
		mockResponse   *models.TransactionResponse
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:          "Successful retrieval of payment status",
			transactionID: "tx123",
			mockResponse: &models.TransactionResponse{
				TransactionRequests: []models.TransactionRequest{
					{
						ID:     "tx123",
						Status: "completed",
						Amount: 100.0,
					},
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "completed",
		},
		{
			name:          "Transaction not found",
			transactionID: "unknown123",
			mockResponse: &models.TransactionResponse{
				TransactionRequests: []models.TransactionRequest{},
			},
			mockError:      nil,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Transaction not found",
		},
		{
			name:           "Payd API error",
			transactionID:  "tx123",
			mockResponse:   nil,
			mockError:      errors.New("Payd API error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to get transaction requests",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient.On("GetTransactionRequests", cfg.AccountID, cfg.APIUsername, cfg.APIPassword).Return(tc.mockResponse, tc.mockError).Once()

			req, _ := http.NewRequest(http.MethodGet, "/payments/status/"+tc.transactionID, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}
