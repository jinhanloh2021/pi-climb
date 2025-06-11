package initialisers

import "github.com/jinhanloh2021/beta-blocker/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
