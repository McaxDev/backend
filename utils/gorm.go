package utils

import "gorm.io/gorm"

func LoadUserBaseInfo(db *gorm.DB) *gorm.DB {
	return db.Select("name", "exp").Preload("Avatar", func(db *gorm.DB) *gorm.DB {
		return db.Select("filename")
	})
}
