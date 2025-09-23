package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/middleware"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
	"github.com/jinhanloh2021/pi-climb/internal/service"
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
	if errors.Is(err, repository.ErrFollowNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Follow to %s not found or not accessible", body.UserID)})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error following user %s", body.UserID)})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *FollowHandler) GetFollowers(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	var targetUserID uuid.UUID = userID
	if paramID := c.Param("id"); paramID != "" {
		if parsedID, err := uuid.Parse(paramID); err == nil {
			targetUserID = parsedID
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse param uuid"})
			return
		}
	}

	followers, err := h.followService.GetFollowers(c, userID, targetUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get followers"})
		return
	}
	c.JSON(http.StatusOK, followers)
}

func (h *FollowHandler) GetFollowing(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	var targetUserID uuid.UUID = userID
	if paramID := c.Param("id"); paramID != "" {
		if parsedID, err := uuid.Parse(paramID); err == nil {
			targetUserID = parsedID
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse param uuid"})
			return
		}
	}

	following, err := h.followService.GetFollowing(c, userID, targetUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get following"})
		return
	}
	c.JSON(http.StatusOK, following)
}

// Get if loggged in user following/follows_you? to target_user
func (h *FollowHandler) GetFollowRelationship(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	var targetUserID uuid.UUID = userID
	if paramID := c.Param("id"); paramID != "" {
		if parsedID, err := uuid.Parse(paramID); err == nil {
			targetUserID = parsedID
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse param uuid"})
			return
		}
	}
	userToTarget, targetToUser, err := h.followService.GetFollowRelationship(c, userID, targetUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get relationship with user %s", targetUserID.String())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"following": userToTarget != nil, "follows_you": targetToUser != nil})
	return
}
