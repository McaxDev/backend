package dbs

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content string `json:"content" gorm:"comment:'题目内容'"`
	Class   string `json:"class" gorm:"comment:'题目分类'"`
	Answer  rune   `json:"answer" gorm:"comment:'正确答案'"`
}
