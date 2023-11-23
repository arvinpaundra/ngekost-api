package entity

import (
	"time"

	"gorm.io/gorm"
)

type Common struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
