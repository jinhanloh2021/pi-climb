package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) GetMyUser(c *gin.Context) {
	userID, _ := middleware.GetUserID(c) // caller is also target, get self
	user, err := h.userService.GetUserByID(c.Request.Context(), userID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving User"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	user, err := h.userService.GetUserByUsername(c.Request.Context(), c.Param("username"), userID) // url param
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", c.Param("username"))})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user by username"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	var reqBody dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), userID, &reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error updating user %s", userID)})
		return
	}
	c.JSON(http.StatusOK, user)
}
