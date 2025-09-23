package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/config"
)

// CustomClaims defines the claims expected in Supabase JWTs
type CustomClaims struct {
	Aud          string         `json:"aud"`
	Exp          int64          `json:"exp"`
	Iat          int64          `json:"iat"`
	Iss          string         `json:"iss"`
	Sub          string         `json:"sub"`
	Email        string         `json:"email"`
	UserMetaData map[string]any `json:"user_metadata"`
	AppMetaData  map[string]any `json:"app_metadata"`
	jwt.RegisteredClaims
}

type JWTValidator interface {
	ValidateSupabaseJWT(tokenString string) (uuid.UUID, *CustomClaims, error)
}

type supabaseJWTValidator struct{}

func NewSupabaseJWTValidator() JWTValidator {
	return &supabaseJWTValidator{}
}

// Validates the JWT issued by Supabase
func (s *supabaseJWTValidator) ValidateSupabaseJWT(tokenString string) (uuid.UUID, *CustomClaims, error) {
	// Supabase JWTs are signed with HMAC SHA256 using the JWT secret from your project settings.
	jwtSecret := config.LoadConfig().SupabaseJWTSecret

	if jwtSecret == "" {
		return uuid.Nil, nil, errors.New("SUPABASE_JWT_SECRET not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return uuid.Nil, nil, errors.New("invalid token claims")
	}

	// Basic claims validation, Supabase handles most of this
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return uuid.Nil, nil, errors.New("token expired")
	}

	userID, err := uuid.Parse(claims.Sub)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	return userID, claims, nil
}
