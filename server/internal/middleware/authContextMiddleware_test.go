package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	return c, w
}

func HelperSetUserUUIDMiddleware(testUUID any) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(UserIDKey, testUUID)
		c.Next()
	}
}

func TestGetUserUUID_Success(t *testing.T) {
	c, _ := createTestContext()
	expectedUUID := uuid.New()
	c.Set(UserIDKey, expectedUUID)

	actualUUID, ok := GetUserID(c)
	assert.True(t, ok, "Expected UUID to be found")
	assert.Equal(t, expectedUUID, actualUUID, "Retrieved UUID should match expected UUID")
}

func TestGetUserUUID_NotFound(t *testing.T) {
	c, _ := createTestContext()
	// UUID is intentionally NOT set in the context
	// expectedUUID := uuid.New()
	// c.Set(UserUUIDKey, expectedUUID)

	actualUUID, ok := GetUserID(c)
	assert.False(t, ok, "Expected UUID not to be found")
	assert.Equal(t, uuid.Nil, actualUUID, "Retrieved UUID should be nil UUID")
}

func TestGetUserUUID_InvalidType(t *testing.T) {
	c, _ := createTestContext()
	c.Set(UserIDKey, "string not UUID type")

	actualUUID, ok := GetUserID(c)
	assert.False(t, ok, "Expected UUID not to be found due to invalid type")
	assert.Equal(t, uuid.Nil, actualUUID, "Retrieved UUID should be nil UUID for invalid type")
}

func TestUserAuthContextMiddleware_Success(t *testing.T) {
	// Arrange
	testUUID := uuid.New()
	r := gin.New()
	r.Use(HelperSetUserUUIDMiddleware(testUUID)).Use(UserAuthContextMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 OK")
	assert.Contains(t, w.Body.String(), "Access granted", "Expected access granted message")
}

func TestUserAuthContextMiddleware_MissingUUID(t *testing.T) {
	r := gin.New()
	// r.use(HelperSetUserUUIDMiddleware(testUUID))
	r.Use(UserAuthContextMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"}) // This should NOT be reached
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error")
	assert.Contains(t, w.Body.String(), "Authentication context missing or invalid, user UUID not found in context after auth")
}

func TestUserAuthContextMiddleware_InvalidUUIDType(t *testing.T) {
	r := gin.New()
	testUUID := "Invalid string UUID"
	r.Use(HelperSetUserUUIDMiddleware(testUUID)).Use(UserAuthContextMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"}) // This should NOT be reached
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected HTTP 500 Internal Server Error")
	assert.Contains(t, w.Body.String(), "Authentication context missing or invalid, user UUID not found in context after auth")
}
