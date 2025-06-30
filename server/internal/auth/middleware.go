package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	UserUUIDKey = "SupabaseID"
	ClaimsKey   = "Claims"
)

// Validates the JWT from the Authorization header or cookie and sets the userID and claims in the Gin context
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

		userUUID, claims, err := ValidateSupabaseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			return
		}

		c.Set(UserUUIDKey, userUUID)
		c.Set(ClaimsKey, claims)

		c.Next() // Call Next handler
	}
}

// Get requesting user's UUID
func GetUserUUID(c *gin.Context) (uuid.UUID, bool) {
	userUUIDAny, ok := c.Get(UserUUIDKey)
	if !ok {
		return uuid.Nil, false
	}
	userUUID, ok := userUUIDAny.(uuid.UUID)
	if !ok {
		return uuid.Nil, false
	}
	return userUUID, true
}

// UserAuthContextMiddleware ensures the user UUID is present and valid in the context, run after main auth middleware
func UserAuthContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := GetUserUUID(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication context missing or invalid: User UUID not found in context after auth"})
			c.Abort()
			return
		}
		c.Next()
	}
}
