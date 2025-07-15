package seed

import (
	"log"

	"github.com/jinhanloh2021/beta-blocker/internal/models"
	"github.com/jinhanloh2021/beta-blocker/scripts/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedPosts() {
	seedConfig := config.LoadSeedConfig()
	dsn := seedConfig.DbURLPostgresRole
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database as postgres role")
	}

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
