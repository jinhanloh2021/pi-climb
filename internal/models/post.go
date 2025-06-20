package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID           uint    `gorm:"not null;index"`
	User             *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Caption          *string `gorm:"size:512"`
	HoldColour       *string `gorm:"size:64"`
	DifficultyRating *string `gorm:"size:64"`

	Media []Media `gorm:"polymorphic:Owner;polymorphicValue:posts;constraint:OnDelete:CASCADE;"`

	Likes     []Like `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	LikeCount uint   `gorm:"not null;default:0"` // update in app level in hook, AfterUpdate

	CommentCount uint      `gorm:"not null;default:0"`
	Comments     []Comment `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	GymID *uint // optional Gym relation
	Gym   *Gym  `gorm:"foreignKey:GymID;references:ID"`
}
