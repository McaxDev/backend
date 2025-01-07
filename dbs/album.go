package dbs

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Cover  string `json:"cover" gorm:"comment:'封面文件名'"`
	Order  int    `json:"order" gorm:"comment:'排序'"`
	Title  string `json:"name" gorm:"unique;comment:'标题'"`
	Folder string `json:"folder" gorm:"unique;comment:'目录名'"`
	UserID uint   `json:"userId" gorm:"comment:'创建者'"`
	User   User
	Images []Image
}
