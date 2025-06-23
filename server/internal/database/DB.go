package database

import (
	"fmt"
	"log"

	"github.com/jinhanloh2021/beta-blocker/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	cfg := config.LoadConfig() // Load config
	dsn := cfg.DatabaseURL

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to the database")
}
