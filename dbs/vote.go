package dbs

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	IssueID uint   `json:"issueId" gorm:"not null;comment:议题ID"`
	UserID  uint   `json:"userId" gorm:"not null;comment:用户ID"`
	User    User   `gorm:"constraint:OnDelete:CASCADE"`
	Agree   bool   `json:"agree" gorm:"not null;comment:赞成"`
	Content string `json:"content" gorm:"type:TEXT;comment:意见"`
}
