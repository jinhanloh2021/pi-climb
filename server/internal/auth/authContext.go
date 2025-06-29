package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
