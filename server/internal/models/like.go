package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	UserID uuid.UUID `gorm:"primaryKey;not null" json:"user_id"`
	User   *User     `gorm:"foreignKey:UserID;references:ID" json:"user"`

	PostID uint  `gorm:"primaryKey;not null;index" json:"post_id"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID" json:"post"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"` // to re-like, find soft-deleted like and set DeletedAt to NULL
}
