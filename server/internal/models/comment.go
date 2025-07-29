package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Text string `gorm:"size:512;not null" json:"text"`

	UserID uuid.UUID `gorm:"not null" json:"user_id"`
	User   *User     `gorm:"foreignKey:UserID;references:ID" json:"user"`

	PostID uint  `gorm:"not null;index" json:"post_id"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID" json:"post"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
