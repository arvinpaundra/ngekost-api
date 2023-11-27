package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type gormScope func(*gorm.DB) *gorm.DB

func Search(col, search string) gormScope {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			lowerCol := fmt.Sprintf("LOWER(%s)", col)
			lowerSearch := strings.ToLower(search)
			return db.Where(lowerCol+" LIKE ?", "%"+lowerSearch+"%")
		}
		return db
	}
}

func Exact(col, val string) gormScope {
	return func(db *gorm.DB) *gorm.DB {
		if val != "" {
			return db.Where(col+" = ?", val)
		}
		return db
	}
}
