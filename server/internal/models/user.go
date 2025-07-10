package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey"`
	Email       string    `gorm:"unique;not null"`
	Username    string    `gorm:"unique;size:64;not null"`
	Bio         *string   `gorm:"size:255"`
	IsPublic    bool      `gorm:"default:true"`
	PhoneNumber *string   `gorm:"unique"`
	DateOfBirth *time.Time

	Followers      []Follow `gorm:"foreignKey:ToUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	FollowerCount  uint     `gorm:"not null;default:0"`
	Following      []Follow `gorm:"foreignKey:FromUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	FollowingCount uint     `gorm:"not null;default:0"`

	Avatar *Media `gorm:"polymorphic:Owner;polymorphicValue:users;constraint:OnDelete:CASCADE"`

	Posts []Post `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Likes []Like `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Comments []Comment `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Media []Media `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
