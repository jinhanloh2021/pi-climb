package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/dto"
	"github.com/jinhanloh2021/beta-blocker/internal/middleware"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{postService: s}
}

func (h *PostHandler) CreateNewPost(c *gin.Context) {
	var body dto.CreatePostRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	userID, _ := middleware.GetUserID(c)

	post, err := h.postService.CreatePost(c.Request.Context(), userID, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error creating post")})
		return
	}
	c.JSON(http.StatusCreated, post)
	return
}
