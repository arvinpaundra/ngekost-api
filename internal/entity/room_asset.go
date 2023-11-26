package entity

import (
	"time"

	"gorm.io/gorm"
)

type RoomAsset struct {
	ID        string
	RoomId    string
	Url       string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
