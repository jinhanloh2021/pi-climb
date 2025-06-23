package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text string `gorm:"size:512;not null"`

	UserID uint  `gorm:"primaryKey;not null"`
	User   *User `gorm:"foreignKey:UserID;references:ID"`

	PostID uint  `gorm:"primaryKey;not null;index"`
	Post   *Post `gorm:"foreignKey:PostID;references:ID"`
}
