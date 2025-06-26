package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/auth"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
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
	claims, ok := auth.GetJwtClaimsFromContext(c) // get logged in user from auth middleware
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT claims not found in context"})
		return
	}
	supabaseID, err := uuid.Parse(claims.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse subject into UUID"})
		return
	}
	user, err := h.userService.GetUserByUUID(c, supabaseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWTclaims": claims, "MyInfo": user})
}

func (h *UserHandler) TrySetDifferentUserDOB(c *gin.Context) {
	callerID, ok := auth.GetUserUUIDFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user ID"})
		return
	}

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
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
