package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/middleware"
	"github.com/jinhanloh2021/pi-climb/internal/service"
)

type FeedHandler struct {
	feedService service.FeedService
}

func NewFeedHandler(s service.FeedService) *FeedHandler {
	return &FeedHandler{feedService: s}
}

func (h *FeedHandler) GetFeed(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	followingCursor := c.Query("following-cursor")
	trendingCursor := c.Query("trending-cursor")
	limitParam := c.Query("limit")
	if limitParam == "" {
		limitParam = "10"
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing limit %s", limitParam)})
		return
	}

	var feedCursor dto.FeedCursor = dto.FeedCursor{FollowingCursor: followingCursor, TrendingCursor: trendingCursor}

	feed, nextCursor, err := h.feedService.GetFeed(c.Request.Context(), userID, &feedCursor, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error getting feed")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": feed, "next_cursor": nextCursor})
	return
}
