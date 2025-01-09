package dbs

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Filename    string `json:"filename" gorm:"not null;unique;type:VARCHAR(255);unique;comment:文件名"`
	Title       string `json:"title" gorm:"not null;unique;type:VARCHAR(255);unique;comment:标题"`
	Description string `json:"description" gorm:"type:TEXT;comment:简介"`
	Order       int    `json:"order" gorm:"comment:排序"`
	Likes       uint   `json:"likes" gorm:"not null;comment:点赞"`
	UserID      uint   `json:"userId" gorm:"comment:上传者"`
	User        User   `gorm:"constraint:OnDelete:SET NULL"`
	AlbumID     uint   `json:"albumId" gorm:"not null;comment:相册ID"`
	Album       Album  `gorm:"constraint:OnDelete:CASCADE"`
}
