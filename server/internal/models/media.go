package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaType string

type Media struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	URL           string `gorm:"not null;size:512" json:"url"`
	StoragePath   string `gorm:"size:512" json:"storage_path"`
	ThumbnailURL  string `gorm:"size:512" json:"thumbnail_url"`
	CompressedURL string `gorm:"size:512" json:"compressed_url"`

	Filename string `gorm:"size:255" json:"filename"`
	FileSize int64  `gorm:"not null" json:"file_size"`

	MimeType string `gorm:"size:100" json:"mime_type"`
	Order    *int   `gorm:"default:0" json:"order"`

	Width    *int `gorm:"default:null" json:"width"`
	Height   *int `gorm:"default:null" json:"height"`
	Duration *int `gorm:"default:null" json:"duration"`

	// Polymorphic Association
	OwnerID   uint   `gorm:"not null;index:idx_owner" json:"owner_id"`
	OwnerType string `gorm:"not null;index:idx_owner" json:"owner_type"`

	// Ownership
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`
}
