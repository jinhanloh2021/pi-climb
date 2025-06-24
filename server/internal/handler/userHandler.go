package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/beta-blocker/internal/auth"
	"github.com/jinhanloh2021/beta-blocker/internal/service"
)

// Handles user-related HTTP requests
type UserHandler struct {
	userService service.UserService
}

// Creates a new UserHandler
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

// GetAuthenticatedUser handles GET /authenticated route
// It requires AuthMiddleware and returns the user's details.
func (h *UserHandler) GetAuthenticatedUser(c *gin.Context) {
	claims, ok := auth.GetJwtClaimsFromContext(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT claims not found in context"})
		return
	}

	// // Check your application's `users` table for SupabaseID.
	// // If it doesn't exist, then create the user's profile.
	// appUser, err := h.userService.GetOrCreateUserBySupabaseID(c.Request.Context(), userID, userEmail)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get or create user profile", "details": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"claims": claims})
}

// New Post handler (example, not fully implemented)
func CreatePost(c *gin.Context) {
	userID, ok := auth.GetUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// Logic to create a post using userID
	c.JSON(http.StatusOK, gin.H{"message": "Post created by user " + userID.String()})
}
