package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL       string
	AnonKey           string
	DbURLPostgresRole string
}

var (
	seedConfig *Config
	once       sync.Once
)

func LoadSeedConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		supabaseURL := os.Getenv("LOCAL_SUPABASE_URL")
		anonKey := os.Getenv("LOCAL_ANON_KEY")
		dbURLPostgresRole := os.Getenv("LOCAL_POSTGRES_DATABASE_URL")

		seedConfig = &Config{
			SupabaseURL:       supabaseURL,
			AnonKey:           anonKey,
			DbURLPostgresRole: dbURLPostgresRole,
		}

		if seedConfig.SupabaseURL == "" {
			log.Fatal("LOCAL_SUPABASE_URL not set in environment")
		}
		if seedConfig.AnonKey == "" {
			log.Fatal("LOCAL_ANON_KEY not set in environment")
		}
		if seedConfig.DbURLPostgresRole == "" {
			log.Fatal("LOCAL_POSTGRES_DATABASE_URL not set in environment")
		}
	})
	return seedConfig
}
