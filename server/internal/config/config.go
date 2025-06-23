package config

import (
	"fmt"
	"log"
	"os"
	"sync" // For singleton pattern if needed

	"github.com/joho/godotenv"
)

// Config holds all application configurations
type Config struct {
	DatabaseURL       string
	SupabaseJWTSecret string // Supabase JWT secret (for verification)
}

var (
	appConfig *Config
	once      sync.Once
)

// LoadConfig loads environment variables and returns the config struct.
// It uses a singleton pattern to ensure config is loaded only once.
func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			// log.Println("No .env file found, assuming environment variables are set.")
			log.Fatal("Error loading .env file") // Only fatal if .env is strictly required
		}
		appEnv := os.Getenv("APP_ENV")
		dbURL := os.Getenv(fmt.Sprintf("%s_DATABASE_URL", appEnv))
		jwtSecret := os.Getenv(fmt.Sprintf("%s_JWT_SECRET", appEnv))

		appConfig = &Config{
			DatabaseURL:       dbURL,
			SupabaseJWTSecret: jwtSecret,
		}

		if appConfig.DatabaseURL == "" {
			log.Fatal("DATABASE_URL not set in environment")
		}
		if appConfig.SupabaseJWTSecret == "" {
			log.Fatal("SUPABASE_JWT_SECRET not set in environment")
		}
	})
	return appConfig
}
