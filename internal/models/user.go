package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model includes ID (uint), CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	SupabaseID  uuid.UUID `gorm:"type:uuid;unique;not null;index"`
	Email       string    `gorm:"unique;not null"`
	Username    string    `gorm:"unique;size:64;not null"`
	AvatarURL   *string
	Bio         *string `gorm:"size:255"`
	IsPublic    bool    `gorm:"default:true"`
	PhoneNumber *string `gorm:"unique"`
	DateOfBirth *time.Time
}
