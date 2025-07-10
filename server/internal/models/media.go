package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaType string

type Media struct {
	gorm.Model

	URL           string `gorm:"not null;size:512"`
	StoragePath   string `gorm:"size:512"` // Original storage path
	ThumbnailURL  string `gorm:"size:512"` // Thumbnail/preview URL
	CompressedURL string `gorm:"size:512"` // Compressed version URL

	Filename string `gorm:"size:255"` // Original filename
	FileSize int64  `gorm:"not null"` // Size in bytes
	MimeType string `gorm:"size:100"` // image/jpeg, video/mp4
	Order    *int   `gorm:"default:0"`

	Width    *int `gorm:"default:null"` // Pixel width
	Height   *int `gorm:"default:null"` // Pixel height
	Duration *int `gorm:"default:null"` // Video duration in seconds

	// Polymorphic Association
	OwnerID   uint   `gorm:"not null;index:idx_owner"`
	OwnerType string `gorm:"not null;index:idx_owner"` // post, user, gym,etc.
	// Ownership
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	User   User      `gorm:"foreignKey:UserID"`
}
