package gormdb

import "gorm.io/gorm"

func Load(db *gorm.DB) {
	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Conversion{},
	)
}
