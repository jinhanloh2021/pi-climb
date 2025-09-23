package config

import (
	"log"
	"os"
	"sync"
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
		dbURL := os.Getenv("DATABASE_URL")
		jwtSecret := os.Getenv("JWT_SECRET")

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
