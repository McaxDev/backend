package database

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Metadata
	Wiki []Wiki `json:"-"`
}

type Wiki struct {
	gorm.Model
	Metadata
	Content    string `json:"content" gorm:"comment:'内容'"`
	CategoryID uint   `json:"category_id" gorm:"comment:'分类ID'"`
}

type Metadata struct {
	Path  string `json:"path" gorm:"comment:'路径'"`
	Title string `json:"title" gorm:"comment:'标题'"`
	Order string
}
