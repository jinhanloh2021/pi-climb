package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/config"
)

// CustomClaims defines the claims expected in Supabase JWTs
type CustomClaims struct {
	Aud          string         `json:"aud"`
	Exp          int64          `json:"exp"`
	Iat          int64          `json:"iat"`
	Iss          string         `json:"iss"`
	Sub          string         `json:"sub"` // Supabase User ID (UUID string)
	Email        string         `json:"email"`
	UserMetaData map[string]any `json:"user_metadata"`
	AppMetaData  map[string]any `json:"app_metadata"`
	jwt.RegisteredClaims
}

// Validates the JWT issued by Supabase
func ValidateSupabaseJWT(tokenString string) (uuid.UUID, string, error) {
	// Supabase JWTs are signed with HMAC SHA256 using the JWT secret from your project settings.
	jwtSecret := config.LoadConfig().SupabaseJWTSecret

	if jwtSecret == "" {
		return uuid.Nil, "", errors.New("SUPABASE_JWT_SECRET not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return uuid.Nil, "", fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return uuid.Nil, "", errors.New("invalid token claims")
	}

	// Basic claims validation, Supabase handles most of this
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return uuid.Nil, "", errors.New("token expired")
	}

	userID, err := uuid.Parse(claims.Sub)
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("invalid user ID in token: %w", err)
	}

	return userID, claims.Email, nil
}
