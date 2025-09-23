package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/pi-climb/internal/dto"
	"github.com/jinhanloh2021/pi-climb/internal/middleware"
	"github.com/jinhanloh2021/pi-climb/internal/repository"
	"github.com/jinhanloh2021/pi-climb/internal/service"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: s}
}

func (h *CommentHandler) GetComments(c *gin.Context) {
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

	comments, err := h.commentService.GetComments(c, postID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get comments for post %d", postID)})
		return
	}
	c.JSON(http.StatusOK, comments)
	return
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
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

	var body dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	comment, err := h.commentService.CreateComment(c, postID, &body, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		return
	}
	c.JSON(http.StatusCreated, comment)
	return
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var commentID uint
	if paramID := c.Param("id"); paramID != "" {
		if parsedID, err := strconv.ParseUint(paramID, 10, 32); err == nil {
			commentID = uint(parsedID)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse param uint"})
			return
		}
	}
	err := h.commentService.DeleteComment(c, commentID, userID)
	if errors.Is(err, repository.ErrCommentNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Comment %d not found or not accessible", commentID)})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete comment %d", commentID)})
		return
	}
	c.Status(http.StatusNoContent)
}
