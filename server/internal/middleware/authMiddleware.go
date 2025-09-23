package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinhanloh2021/pi-climb/internal/auth"
)

const (
	UserIDKey = "UserID"
	ClaimsKey = "Claims"
)

// Validates the JWT from the Authorization header or cookie and sets the userID and claims in the Gin context
func AuthMiddleware(validator auth.JWTValidator) gin.HandlerFunc {
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

		userID, claims, err := validator.ValidateSupabaseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			return
		}

		c.Set(UserIDKey, userID)
		c.Set(ClaimsKey, claims)

		c.Next() // Call Next handler
	}
}
