package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAlbum(c *gin.Context, user *dbs.User, req struct {
	Folder string
	Title  string
	UserID uint
}) {

	if err := DB.Where(
		"folder = ? OR title = ?", req.Folder, req.Title,
	).First(new(dbs.Album)).Error; err == nil {
		c.JSON(400, utils.Resp("已存在同名相册", nil, nil))
		return
	}

	if err := user.ExecWithCoins(
		DB, 30, false, func(tx *gorm.DB) error {
			return DB.Model(new(dbs.Album)).Create(&req).Error
		},
	); err != nil {
		c.JSON(500, utils.Resp("创建相册失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("创建相册成功", nil, nil))
}
