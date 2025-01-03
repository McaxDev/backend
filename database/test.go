package database

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content string `gorm:"comment:'题目内容'"`
	OptionA string `gorm:"comment:'A项内容'"`
	OptionB string `gorm:"comment:'B项内容'"`
	OptionC string `gorm:"comment:'C项内容'"`
	OptionD string `gorm:"comment:'D项内容'"`
	Answer  rune   `gorm:"comment:'正确答案'"`
}
