package models

import (
	"time"

	"gorm.io/gorm"
)

type Gym struct {
	ID            uint     `gorm:"primarykey" json:"id"`
	Name          string   `gorm:"not null;index" json:"name"`
	GradingSystem []string `gorm:"type:text[];not null" json:"grading_system"`

	// Derived from Google Place Details // https://developers.google.com/maps/documentation/places/web-service/place-details
	GooglePlaceID *string  `gorm:"unique" json:"google_place_id"`
	GoogleMapsURI *string  `json:"google_maps_uri"`
	Address       *string  `json:"address"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`

	Posts []Post `gorm:"foreignKey:GymID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"posts"`

	Media []Media `gorm:"polymorphic:Owner;polymorphicValue:gyms;constraint:OnDelete:CASCADE;" json:"media"`

	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
