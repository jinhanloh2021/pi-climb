package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/repository"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
)

type LikeHandler struct {
	likeService service.LikeService
}

func NewLikeHandler(s service.LikeService) *LikeHandler {
	return &LikeHandler{likeService: s}
}

func (h *LikeHandler) CreateLike(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var postID uint
	if paramID := c.Param("id"); paramID != "" {
		if parsedID, err := strconv.ParseUint(paramID, 10, 32); err == nil {
			postID = uint(parsedID)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse param uint"})
			return
		}
	}

	like, err := h.likeService.CreateLike(c, userID, postID)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyLiked) {
			c.JSON(http.StatusConflict, gin.H{"error": "You are already liked this post"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error likeing post %s", postID)})
		return
	}
	c.JSON(http.StatusCreated, like)
	return
}
