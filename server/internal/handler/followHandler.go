package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
)

type FollowHandler struct {
	followService service.FollowService
}

func NewFollowHandler(s service.FollowService) *FollowHandler {
	return &FollowHandler{followService: s}
}

func (h *FollowHandler) CreateFollow(c *gin.Context) {
	fromUserID, _ := middleware.GetUserID(c)
	var body dto.FollowRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	toUserID, err := uuid.Parse(body.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing user UUID %s", body.UserID)})
		return
	}
	follow, err := h.followService.CreateFollow(c, fromUserID, toUserID)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyFollowing) {
			c.JSON(http.StatusConflict, gin.H{"error": "You are already following this user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error following user %s", body.UserID)})
		return
	}

	c.JSON(http.StatusCreated, follow)
}

func (h *FollowHandler) DeleteFollow(c *gin.Context) {
	fromUserID, _ := middleware.GetUserID(c)
	var body dto.FollowRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	toUserID, err := uuid.Parse(body.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing user UUID %s", body.UserID)})
		return
	}
	err = h.followService.DeleteFollow(c, fromUserID, toUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error following user %s", body.UserID)})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *FollowHandler) GetFollowers(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	followers, err := h.followService.GetFollowers(c, userID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get followers"})
		return
	}
	c.JSON(http.StatusOK, followers)
}

func (h *FollowHandler) GetFollowing(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	following, err := h.followService.GetFollowing(c, userID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get following"})
		return
	}
	c.JSON(http.StatusOK, following)
}
