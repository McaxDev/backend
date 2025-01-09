package dbs

import "gorm.io/gorm"

type Wiki struct {
	gorm.Model
	Path     string `json:"path" gorm:"not null;type:VARCHAR(255);comment:路径"`
	Title    string `json:"title" gorm:"not null;type:VARCHAR(255);comment:标题"`
	Order    int    `json:"order" gorm:"comment:次序"`
	Markdown string `json:"content" gorm:"type:TEXT;comment:内容"`
	HTML     string `json:"html" gorm:"type:TEXT;comment:HTML内容"`
	Category string `json:"categoryId" gorm:"type:VARCHAR(255);comment:分类ID"`
}
