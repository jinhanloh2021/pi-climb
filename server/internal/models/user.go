package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primarykey" json:"id"`
	Email       string     `gorm:"unique;not null" json:"email"`
	Username    *string    `gorm:"unique;size:64" json:"username"`
	Bio         *string    `gorm:"size:256" json:"bio"`
	IsPublic    bool       `gorm:"default:true" json:"is_public"`
	PhoneNumber *string    `gorm:"unique" json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	// Relationships - be careful about circular references
	Followers      []Follow `gorm:"foreignKey:ToUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"followers,omitempty"`
	FollowerCount  uint     `gorm:"not null;default:0" json:"follower_count"`
	Following      []Follow `gorm:"foreignKey:FromUserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"following,omitempty"`
	FollowingCount uint     `gorm:"not null;default:0" json:"following_count"`

	Avatar   *Media    `gorm:"polymorphic:Owner;polymorphicValue:users;constraint:OnDelete:CASCADE" json:"avatar,omitempty"`
	Posts    []Post    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"posts,omitempty"`
	Likes    []Like    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"likes,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments,omitempty"`
	Media    []Media   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE" json:"media,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
