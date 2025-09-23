package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	SupabaseURL       string
	AnonKey           string
	DbURLPostgresRole string
	ServiceRoleKey    string
}

var (
	seedConfig *Config
	once       sync.Once
)

func LoadSeedConfig() *Config {
	once.Do(func() {
		supabaseURL := os.Getenv("SUPABASE_URL")
		anonKey := os.Getenv("ANON_KEY")
		dbURLPostgresRole := os.Getenv("POSTGRES_DATABASE_URL")
		serviceRoleKey := os.Getenv("SERVICE_ROLE_KEY")

		seedConfig = &Config{
			SupabaseURL:       supabaseURL,
			AnonKey:           anonKey,
			DbURLPostgresRole: dbURLPostgresRole, // only for seed
			ServiceRoleKey:    serviceRoleKey,
		}

		if seedConfig.SupabaseURL == "" {
			log.Fatal("SUPABASE_URL not set in environment")
		}
		if seedConfig.AnonKey == "" {
			log.Fatal("ANON_KEY not set in environment")
		}
		if seedConfig.DbURLPostgresRole == "" {
			log.Fatal("POSTGRES_DATABASE_URL not set in environment")
		}
		if seedConfig.ServiceRoleKey == "" {
			log.Fatal("SERVICE_ROLE_KEY not set in environment")
		}
	})
	return seedConfig
}
