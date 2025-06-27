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

		userUUID, claims, err := ValidateSupabaseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			return
		}

		// Set user ID and email for subsequent handlers
		c.Set("userUUID", userUUID)
		c.Set("claims", claims)

		c.Next() // Call Next handler
	}
}

// TODO: Remove this, can get claims and userUUID directly from context
func GetJwtClaimsFromContext(c *gin.Context) (*CustomClaims, bool) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, false
	}
	return claims.(*CustomClaims), true
}

func GetUserUUIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	claims, ok := GetJwtClaimsFromContext(c)
	if !ok {
		return uuid.Nil, false
	}
	id, err := uuid.Parse(claims.Sub)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}
