package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Techbite-sudo/payd-payment-polling-service/auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestRegister(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	handler := NewHandler(db, "test-secret")

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", handler.Register)

	testCases := []struct {
		name           string
		requestBody    models.RegisterRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful registration",
			requestBody: models.RegisterRequest{
				Username: "testuser",
				Password: "testpassword",
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "User registered successfully",
		},
		{
			name: "Duplicate username",
			requestBody: models.RegisterRequest{
				Username: "testuser",
				Password: "testpassword",
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to create user",
		},
		{
			name: "Missing username",
			requestBody: models.RegisterRequest{
				Password: "testpassword",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body",
		},
		{
			name: "Missing password",
			requestBody: models.RegisterRequest{
				Username: "testuser",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedBody)
		})
	}
}

func TestLogin(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	handler := NewHandler(db, "test-secret")

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	// Register a user first
	registerReq := models.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	body, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	testCases := []struct {
		name           string
		requestBody    models.LoginRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful login",
			requestBody: models.LoginRequest{
				Username: "testuser",
				Password: "testpassword",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "token",
		},
		{
			name: "Invalid credentials",
			requestBody: models.LoginRequest{
				Username: "testuser",
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Invalid credentials",
		},
		{
			name: "Non-existent user",
			requestBody: models.LoginRequest{
				Username: "nonexistentuser",
				Password: "testpassword",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Invalid credentials",
		},
		{
			name: "Missing username",
			requestBody: models.LoginRequest{
				Password: "testpassword",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body",
		},
		{
			name: "Missing password",
			requestBody: models.LoginRequest{
				Username: "testuser",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedBody)
		})
	}
}
