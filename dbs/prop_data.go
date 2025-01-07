package dbs

import "gorm.io/gorm"

type PropData struct {
	gorm.Model
	PID         string `json:"pid" gorm:"comment:'物品ID'"`
	Name        string `json:"name" gorm:"comment:'名称'"`
	Description string `json:"description" gorm:"comment:'介绍'"`
	Icon        string `json:"icon" gorm:"comment:'图标'"`
}
