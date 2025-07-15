package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL       string
	SupabaseJWTSecret string
}

var (
	appConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
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
