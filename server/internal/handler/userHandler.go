package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/auth"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

type UpdateDOBRequest struct {
	DateOfBirth  time.Time `json:"date_of_birth"`
	TargetUserID uuid.UUID `json:"target_user_id"`
}

func (h *UserHandler) GetMyUser(c *gin.Context) {
	supabaseID, _ := auth.GetUserUUID(c)
	user, err := h.userService.GetUserByUUID(c, supabaseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	userUUID, _ := auth.GetUserUUID(c)
	user, err := h.userService.GetUserByUsername(c, c.Param("username"), userUUID) // url param
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", c.Param("username"))})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user by username"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) TrySetDifferentUserDOB(c *gin.Context) {
	callerID, _ := auth.GetUserUUID(c)
	var reqBody UpdateDOBRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	targetID := reqBody.TargetUserID
	DOB := reqBody.DateOfBirth

	user, err := h.userService.SetUserDOB(c, targetID, callerID, &DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error setting DOB. " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
