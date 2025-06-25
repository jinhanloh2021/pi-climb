package handler

import (
	"net/http"

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
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	c.JSON(http.StatusOK, gin.H{"JWTclaims": claims, "MyInfo": user})
}
