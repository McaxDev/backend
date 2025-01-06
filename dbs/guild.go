package dbs

import "gorm.io/gorm"

type Guild struct {
	gorm.Model
	GID     string `json:"gid" gorm:"comment:'公会ID'"`
	Name    string `json:"name" gorm:"comment:'公会名'"`
	Number  uint   `json:"number" gorm:"comment:'公会人数'"`
	Logo    string `json:"logo" gorm:"comment:'LOGO路径'"`
	Profile string `json:"profile" gorm:"comment:'公会介绍'"`
	Money   uint   `json:"money" gorm:"comment:'公会资金'"`
	User    []User
}
