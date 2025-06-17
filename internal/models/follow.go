package models

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	FromUserID uint  `gorm:"primaryKey;not null;index"`
	ToUserID   uint  `gorm:"primaryKey;not null;index"`
	FromUser   *User `gorm:"foreignKey:FromUserID;references:ID"`
	ToUser     *User `gorm:"foreignKey:ToUserID;references:ID"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}
