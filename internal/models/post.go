package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID           uint    `gorm:"not null;index"`
	User             *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Media            []Media `gorm:"polymorphic:Owner;polymorphicValue:posts;constraint:OnDelete:CASCADE;"`
	Caption          *string `gorm:"size:512"`
	HoldColour       *string `gorm:"size:64"`
	DifficultyRating *string `gorm:"size:64"`
	// reference gym
}
