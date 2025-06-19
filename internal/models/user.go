package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	SupabaseID  uuid.UUID `gorm:"type:uuid;unique;not null;index"`
	Email       string    `gorm:"unique;not null"`
	Username    string    `gorm:"unique;size:64;not null"`
	Bio         *string   `gorm:"size:255"`
	IsPublic    bool      `gorm:"default:true"`
	PhoneNumber *string   `gorm:"unique"`
	DateOfBirth *time.Time

	Followers []Follow `gorm:"foreignKey:ToUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Following []Follow `gorm:"foreignKey:FromUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Avatar *Media `gorm:"polymorphic:Owner;polymorphicValue:users;constraint:OnDelete:CASCADE"`

	Posts []Post `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Likes []Like `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
