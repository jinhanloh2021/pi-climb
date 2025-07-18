package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	StorageKey string `gorm:"size:512" json:"storage_key"` // e.g. user_id/filename.jpg
	Bucket     string `gorm:"size:256" json:"bucket"`

	OriginalName string `gorm:"size:512" json:"original_name"`
	FileSize     int64  `gorm:"not null" json:"file_size"`

	MimeType string `gorm:"size:128" json:"mime_type"`
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
