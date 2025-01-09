package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func EditAlbum(c *gin.Context, user *dbs.User, req struct {
	ID        uint
	Cover     string
	Title     string
	OnlyAdmin bool
	GuildID   uint
	Order     int
}) {

	var album dbs.Album
	if err := DB.Where("id = ?", req.ID).First(
		&album,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取相册失败", err, nil))
		return
	}

	if !CheckEditAlbum(user, &album) {
		c.JSON(400, utils.Resp("你没有权限", nil, nil))
		return
	}

	if err := DB.Model(&album).Updates(&req).Error; err != nil {
		c.JSON(500, utils.Resp("更新相册失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新相册成功", nil, nil))
}
