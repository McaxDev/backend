package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func DelAlbum(c *gin.Context, user *dbs.User, id uint) {

	var album dbs.Album
	if err := DB.Where("id = ?", id).First(
		&album,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取相册失败", err, nil))
		return
	}

	if !CheckEditAlbum(user, &album) {
		c.JSON(400, utils.Resp("你没有权限操作", nil, nil))
		return
	}

	if len(album.Images) != 0 {
		c.JSON(400, utils.Resp("相册不为空，无法删除", nil, nil))
		return
	}

	if err := DB.Delete(&album).Error; err != nil {
		c.JSON(500, utils.Resp("删除相册失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("删除相册成功", nil, nil))
}
