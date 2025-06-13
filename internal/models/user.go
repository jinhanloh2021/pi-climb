package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model includes ID (uint), CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	SupabaseID uuid.UUID `gorm:"type:uuid;unique;not null"`
	Email      string    `gorm:"unique;not null"`
	Username   string    `gorm:"unique;size:50"`
	AvatarURL  string
}
