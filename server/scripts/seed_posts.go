package main

import (
	"log"
	"os"

	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func seedPosts() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("LOCAL_DATABASE_URL_POSTGRES_ROLE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var users []models.User
	db.Find(&users)

	for _, user := range users {
		caption := "abc"
		holdColour := "red"
		grade := "v2"

		post := models.Post{
			Caption:    &caption,
			HoldColour: &holdColour,
			Grade:      &grade,
			UserID:     user.ID,
			Views:      0,
		}
		db.Create(&post)
	}
}
