package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Caption    *string `gorm:"size:512" json:"caption"`
	HoldColour *string `gorm:"size:64" json:"hold_colour"`
	Grade      *string `gorm:"size:64" json:"grade"`

	UserID uuid.UUID `gorm:"not null;index" json:"user_id"`
	User   *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`

	Media []Media `gorm:"polymorphic:Owner;polymorphicValue:posts;constraint:OnDelete:CASCADE;" json:"media,omitempty"`

	Likes     []Like `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"likes,omitempty"`
	LikeCount uint   `gorm:"not null;default:0" json:"like_count"`

	Comments     []Comment `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"comments,omitempty"`
	CommentCount uint      `gorm:"not null;default:0" json:"comment_count"`

	GymID *uint `json:"gym_id"`
	Gym   *Gym  `gorm:"foreignKey:GymID;references:ID" json:"gym,omitempty"`
}
