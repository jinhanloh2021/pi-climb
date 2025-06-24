package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{"JWTclaims": claims})
}
