package models

import (
	"time"

	"gorm.io/gorm"
)

type MediaVersion struct {
	ID uint `gorm:"primarykey" json:"id"`

	StorageKey string `gorm:"size:512;not null" json:"storage_key"` // e.g. user_id/filename.jpg
	Bucket     string `gorm:"size:256;not null" json:"bucket"`

	FileSize    uint64 `gorm:"not null" json:"file_size"`
	VersionType string `gorm:"size:128;not null" json:"version_type"` // thumbnail, mobile

	Width    *uint `gorm:"default:null" json:"width"`
	Height   *uint `gorm:"default:null" json:"height"`
	Duration *uint `gorm:"default:null" json:"duration"`

	MediaID uint   `gorm:"primaryKey;not null;index" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID" json:"media"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
