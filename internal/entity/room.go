package entity

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	ID          string
	KostId      string
	Name        string
	Quantity    int
	Price       float64
	Category    *string
	Description *string
	Image       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Kost        Kost
	Assets      []RoomAsset
}
