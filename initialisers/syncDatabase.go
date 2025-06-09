package initialisers

import "github.com/jinhanloh2021/book-anything/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
