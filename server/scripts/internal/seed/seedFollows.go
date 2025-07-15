package seed

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/scripts/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FollowSeedInput struct {
	FromUserEmail string `json:"from_user_id"`
	ToUserEmail   string `json:"to_user_id"`
}

func SeedFollows() {
	// Load the seed data
	file, err := os.Open("./scripts/data/followsData.json")
	if err != nil {
		log.Fatalf("Failed to open follows data file: %v", err)
	}
	defer file.Close()

	var follows []FollowSeedInput
	if err := json.NewDecoder(file).Decode(&follows); err != nil {
		log.Fatalf("Failed to decode follows data: %v", err)
	}

	// Connect to database as postgres role
	seedConfig := config.LoadSeedConfig()
	dsn := seedConfig.DbURLPostgresRole
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Printf("Seeding %d follows...\n", len(follows))

	for _, f := range follows {
		fromID, err := getUserIDByEmail(db, f.FromUserEmail)
		if err != nil {
			log.Printf("Skipping follow: cannot find from_user %s: %v", f.FromUserEmail, err)
			continue
		}

		toID, err := getUserIDByEmail(db, f.ToUserEmail)
		if err != nil {
			log.Printf("Skipping follow: cannot find to_user %s: %v", f.ToUserEmail, err)
			continue
		}

		follow := models.Follow{
			FromUserID: fromID,
			ToUserID:   toID,
		}

		if err := db.Create(&follow).Error; err != nil {
			log.Printf("Failed to insert follow (%s -> %s): %v", f.FromUserEmail, f.ToUserEmail, err)
		} else {
			fmt.Printf("Created follow: %s -> %s\n", f.FromUserEmail, f.ToUserEmail)
		}
	}
	fmt.Println("Follow seeding completed")
}

func getUserIDByEmail(db *gorm.DB, email string) (uuid.UUID, error) {
	var user models.User
	findResult := db.Select("id").Where("email = ?", email).First(&user)
	if findResult.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to find user by email %s: %w", email, findResult.Error)
	}
	return user.ID, nil
}
