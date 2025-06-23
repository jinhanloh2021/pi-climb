package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Validates the JWT from the Authorization header or cookie and sets the user ID and email in the Gin context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ""

		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// Fallback: Get token from "Authorization" cookie
		if tokenString == "" {
			cookie, err := c.Cookie("Authorization")
			if err == nil {
				tokenString = cookie
			}
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			return
		}

		userID, userEmail, err := ValidateSupabaseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			return
		}

		// Set user ID and email for subsequent handlers
		c.Set("userID", userID)
		c.Set("userEmail", userEmail)

		c.Next() // Call Next handler
	}
}

// Retrieves the user ID from the Gin context
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	userID, ok := c.Get("userID")
	if !ok {
		return uuid.Nil, false
	}
	return userID.(uuid.UUID), true
}

// Retrieves the user email from the Gin context
func GetUserEmailFromContext(c *gin.Context) (string, bool) {
	userEmail, ok := c.Get("userEmail")
	if !ok {
		return "", false
	}
	return userEmail.(string), true
}
