package models

import (
	"gorm.io/gorm"
)

type Gym struct {
	gorm.Model
	Name string `gorm:"not null;index"`

	// Derived from Google Place Details // https://developers.google.com/maps/documentation/places/web-service/place-details
	GooglePlaceID *string `gorm:"unique"`
	GoogleMapsURI *string
	Address       *string
	Latitude      *float64
	Longitude     *float64

	Posts []Post `gorm:"foreignKey:GymID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Media []Media `gorm:"polymorphic:Owner;polymorphicValue:gyms;constraint:OnDelete:CASCADE;"`
}
