package models

import (
	"gorm.io/gorm"
	"time"
)

type Like struct {
	UserID uint  `gorm:"primaryKey;not null"`
	PostID uint  `gorm:"primaryKey;not null;index"`
	User   *User `gorm:"foreignKey:UserID;references:ID"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt // to re-like, find soft-deleted like and set DeletedAt to NULL
}
