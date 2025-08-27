package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Root of media versions
type Media struct {
	ID uint `gorm:"primarykey" json:"id"`

	OriginalName string `gorm:"size:512;not null" json:"original_name"`

	MimeType string `gorm:"size:128" json:"mime_type"`
	Order    *uint  `gorm:"default:0" json:"order"`

	// Polymorphic Association
	OwnerID   uint   `gorm:"not null;index:idx_owner" json:"owner_id"`
	OwnerType string `gorm:"not null;index:idx_owner" json:"owner_type"`

	// Ownership
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User   *User     `gorm:"foreignKey:UserID" json:"user"`

	MediaVersions []MediaVersion `gorm:"foreignKey:MediaID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"media_version,omitempty"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
