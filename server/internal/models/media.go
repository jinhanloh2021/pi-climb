package models

import (
	"gorm.io/gorm"
)

type MediaType string

// These types can own media
const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
)

type Media struct {
	gorm.Model
	URL       string    `gorm:"not null;size:512"`
	MediaType MediaType `gorm:"type:varchar(32);not null;default:'image'"`
	Order     *int      `gorm:"default:0"`

	// Polymorphic Association
	OwnerID   uint   `gorm:"not null;index:idx_owner"`
	OwnerType string `gorm:"not null;index:idx_owner"` // post, user, gym,etc.
}
