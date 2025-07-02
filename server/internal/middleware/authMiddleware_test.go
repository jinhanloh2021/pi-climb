package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jinhanloh2021/beta-blocker/internal/auth"
)

type MockJWTValidator struct {
	mock.Mock
}

func (m *MockJWTValidator) ValidateSupabaseJWT(tokenString string) (uuid.UUID, *auth.CustomClaims, error) {
	var claims *auth.CustomClaims
	var id uuid.UUID

	args := m.Called(tokenString)
	if args.Get(0) != nil {
		id = args.Get(0).(uuid.UUID)
	}
	if args.Get(1) != nil {
		claims = args.Get(1).(*auth.CustomClaims)
	}
	return id, claims, args.Error(2)
}

func TestAuthMiddleware_Success(t *testing.T) {
	// Arrange
	mockJWTValidator := new(MockJWTValidator)

	testUUID := uuid.New()
	testClaims := &auth.CustomClaims{
		Sub:   testUUID.String(),
		Email: "test@example.com",
		Exp:   time.Now().Add(time.Hour).Unix(),
	}
	testToken := "valid.jwt.token"

	// Mock the JWT validator
	mockJWTValidator.On("ValidateSupabaseJWT", testToken).Return(testUUID, testClaims, nil).Once()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware(mockJWTValidator))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Access granted")

	mockJWTValidator.AssertExpectations(t)
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	// Arrange
	mockJWTValidator := new(MockJWTValidator)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware(mockJWTValidator))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+"") // no bearer token
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	mockJWTValidator.AssertNumberOfCalls(t, "ValidateSupabaseJWT", 0)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.Contains(t, w.Body.String(), "Authorization token not provided")

	mockJWTValidator.AssertExpectations(t)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	// Arrange
	mockJWTValidator := new(MockJWTValidator)
	invalidToken := "invalid.jwt.token"

	// Mock the JWT validator
	mockJWTValidator.On("ValidateSupabaseJWT", invalidToken).Return(uuid.Nil, nil, fmt.Errorf("invalid token")).Once()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware(mockJWTValidator))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.Contains(t, w.Body.String(), "Invalid or expired token")

	mockJWTValidator.AssertExpectations(t)
}
