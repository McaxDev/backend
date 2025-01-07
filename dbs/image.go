package dbs

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Filename    string `json:"filename" gorm:"unique;comment:'文件名'"`
	Title       string `json:"title" gorm:"unique;comment:'标题'"`
	Description string `json:"description" gorm:"comment:'简介'"`
	Order       int    `json:"order" gorm:"comment:'排序'"`
	Likes       uint   `json:"likes" gorm:"comment:'点赞'"`
	UserID      uint   `json:"userId" gorm:"comment:'上传者'"`
	User        User
	AlbumID     uint `json:"albumId" gorm:"comment:'相册ID'"`
	Album       Album
}
