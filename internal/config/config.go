package config

import (
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

		appConfig = &Config{
			DatabaseURL:       os.Getenv("DATABASE_URL"),
			SupabaseJWTSecret: os.Getenv("SUPABASE_JWT_SECRET"), // Get this from Supabase project settings
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
