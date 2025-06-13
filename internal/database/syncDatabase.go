package database

import "github.com/jinhanloh2021/beta-blocker/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
