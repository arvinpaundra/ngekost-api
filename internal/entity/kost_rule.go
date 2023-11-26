package entity

import (
	"time"

	"gorm.io/gorm"
)

type KostRule struct {
	ID          string
	KostId      string
	Title       string
	Priority    string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
