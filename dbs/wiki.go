package dbs

import "gorm.io/gorm"

type Wiki struct {
	gorm.Model
	WikiMeta
	Markdown   string `json:"content" gorm:"comment:'内容'"`
	HTML       string `json:"html" gorm:"comment:'HTML内容'"`
	CategoryID uint   `json:"categoryId" gorm:"comment:'分类ID'"`
}
