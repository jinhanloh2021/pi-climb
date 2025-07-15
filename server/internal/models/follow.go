package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	FromUserID uuid.UUID `gorm:"primaryKey;not null;index" json:"from_user_id"`
	FromUser   *User     `gorm:"foreignKey:FromUserID;references:ID" json:"from_user"`

	ToUserID uuid.UUID `gorm:"primaryKey;not null;index" json:"to_user_id"`
	ToUser   *User     `gorm:"foreignKey:ToUserID;references:ID" json:"to_user"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
