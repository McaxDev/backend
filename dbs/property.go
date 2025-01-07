package dbs

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	UserID     uint
	PropertyID uint
	PropData   PropData
	Count      uint
}
