package seed

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinhanloh2021/pi-climb/internal/models"
	"gorm.io/gorm"
)

func GetUserIDByEmail(db *gorm.DB, email string) (uuid.UUID, error) {
	var user models.User
	findResult := db.Select("id").Where("email = ?", email).First(&user)
	if findResult.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to find user by email %s: %w", email, findResult.Error)
	}
	return user.ID, nil
}
