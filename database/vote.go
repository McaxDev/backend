package database

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	UserID  uint   `gorm:"comment:'用户名'"`
	Title   string `gorm:"comment:'议题标题'"`
	Content string `gorm:"comment:'议题内容'"`
	Vote    []Vote `gorm:"foreignKey:IssueID"`
	User    User   `gorm:"foreignKey:ID;references:UserID"`
}

type Vote struct {
	gorm.Model
	IssueID uint   `gorm:"comment:'议题ID'"`
	UserID  uint   `gorm:"comment:'用户ID'"`
	Agree   bool   `gorm:"comment:'赞成'"`
	Content string `gorm:"comment:'意见'"`
}
