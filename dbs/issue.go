package dbs

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	UserID  uint   `json:"userId" gorm:"comment:'用户名'"`
	Title   string `json:"title" gorm:"comment:'议题标题'"`
	Content string `json:"content" gorm:"comment:'议题内容'"`
	Vote    []Vote
	User    User
}
