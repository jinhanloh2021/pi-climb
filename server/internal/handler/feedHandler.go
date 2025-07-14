package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
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
	var feedCursor dto.FeedCursor = dto.FeedCursor{FollowingCursor: followingCursor, TrendingCursor: trendingCursor}

	feed, nextCursor, err := h.feedService.GetFeed(c.Request.Context(), userID, &feedCursor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error getting feed")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": feed, "nextCursor": nextCursor})
	return
}
