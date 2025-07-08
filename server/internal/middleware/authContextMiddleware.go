package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Get requesting user's UUID
func GetUserID(c *gin.Context) (uuid.UUID, bool) {
	userUUIDAny, ok := c.Get(UserIDKey)
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
		_, ok := GetUserID(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Authentication context missing or invalid, user UUID not found in context after auth"})
			return
		}
		c.Next()
	}
}
