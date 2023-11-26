package db

import "gorm.io/gorm"

type gormScope func(*gorm.DB) *gorm.DB

func Search(col, search string) gormScope {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			return db.Where(col+" LIKE ?", "%"+search+"%")
		}
		return db
	}
}
