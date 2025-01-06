package dbs

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	IssueID uint   `json:"issueId" gorm:"comment:'议题ID'"`
	UserID  uint   `json:"userId" gorm:"comment:'用户ID'"`
	Agree   bool   `json:"agree" gorm:"comment:'赞成'"`
	Content string `json:"content" gorm:"comment:'意见'"`
}
