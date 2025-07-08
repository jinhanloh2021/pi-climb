package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	FromUserID uuid.UUID `gorm:"primaryKey;not null;index"`
	FromUser   *User     `gorm:"foreignKey:FromUserID;references:ID"`

	ToUserID uuid.UUID `gorm:"primaryKey;not null;index"`
	ToUser   *User     `gorm:"foreignKey:ToUserID;references:ID"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}
