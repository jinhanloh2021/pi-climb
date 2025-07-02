package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByUsername(c context.Context, username string, userUUID uuid.UUID) (*models.User, error) {
	// Records arguments which were called with GetUserByUsername
	args := m.Called(c, username, userUUID)

	// This block handles the return values based on what was configured with mock.On().
	// args.Get(0) retrieves the first return value (the *models.User).
	// args.Error(1) retrieves the second return value (the error).
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByUUID(c context.Context, supabaseID uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (m *MockUserService) SetUserDOB(c context.Context, targetID uuid.UUID, callerID uuid.UUID, DOB *time.Time) (*models.User, error) {
	return nil, nil
}

// Returns gin context and http response
func createTestContext(method, url string, body any) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	reqBody := new(bytes.Buffer)
	if body != nil {
		json.NewEncoder(reqBody).Encode(body)
	}
	req, _ := http.NewRequest(method, url, reqBody)
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	return c, w
}

func TestGetByUsername_Success(t *testing.T) {
	// Arrange
	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)

	targetUsername := "testUser"
	testUserUUID := uuid.New()

	expectedUser := &models.User{
		SupabaseID: testUserUUID,
		Username:   targetUsername,
		Email:      "testuser@test.com",
		IsPublic:   true,
	}

	// mock return value of GetUserByUsername
	mockUserService.On("GetUserByUsername", mock.Anything, targetUsername, mock.AnythingOfType("uuid.UUID")).Return(expectedUser, nil).Once()

	c, w := createTestContext(http.MethodGet, "/user/"+targetUsername, nil)
	c.Params = gin.Params{
		gin.Param{Key: "username", Value: targetUsername},
	}
	c.Set(middleware.UserUUIDKey, testUserUUID)

	// Act
	userHandler.GetUserByUsername(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "user")
	assert.NotContains(t, responseBody, "error")

	userMap := responseBody["user"].(map[string]any)
	assert.Equal(t, targetUsername, userMap["Username"])
	assert.Equal(t, testUserUUID.String(), userMap["SupabaseID"])

	// verifies each expected method was called with correct arguments and number of times. mock.On()
	mockUserService.AssertExpectations(t)
}

func TestGetByUsername_NotFound(t *testing.T) {
	// Arrange
	mockUserService := new(MockUserService)
	userHandler := NewUserHandler(mockUserService)

	targetUsername := "nonExistentUser"
	testUserUUID := uuid.New()

	mockUserService.On("GetUserByUsername", mock.Anything, targetUsername, mock.AnythingOfType("uuid.UUID")).Return(nil, gorm.ErrRecordNotFound).Once()

	c, w := createTestContext(http.MethodGet, "/user/"+targetUsername, nil)
	c.Params = gin.Params{
		gin.Param{Key: "username", Value: targetUsername},
	}
	c.Set(middleware.UserUUIDKey, testUserUUID)

	// Act
	userHandler.GetUserByUsername(c)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	var responseBody map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "error")
	assert.NotContains(t, responseBody, "user")

	errString := responseBody["error"].(string)
	assert.Equal(t, fmt.Sprintf("User %s not found", targetUsername), errString)

	mockUserService.AssertExpectations(t)
}
