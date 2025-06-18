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
	OwnerID   uint   `gorm:"not null;index:idx_owner"`
	OwnerType string `gorm:"not null;index:idx_owner"`

	URL       string    `gorm:"not null;size:512"`
	MediaType MediaType `gorm:"type:varchar(32);not null;default:'image'"`
	Order     *int      `gorm:"default:0"`
}
