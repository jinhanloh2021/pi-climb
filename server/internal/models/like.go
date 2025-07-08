package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	UserID uuid.UUID `gorm:"primaryKey;not null"`
	User   *User     `gorm:"foreignKey:UserID;references:ID"`

	PostID uint  `gorm:"primaryKey;not null;index"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt // to re-like, find soft-deleted like and set DeletedAt to NULL
}
