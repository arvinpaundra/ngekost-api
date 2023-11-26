package entity

import (
	"time"

	"gorm.io/gorm"
)

type Kost struct {
	ID              string
	OwnerId         string
	Name            string
	Description     string
	Type            string
	PaymentInterval string
	Province        string
	City            string
	District        string
	Subdistrict     string
	Latitude        float64
	Longitude       float64
	Image           *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Owner           Owner
	Rooms           []Room
	Rules           []KostRule
}
