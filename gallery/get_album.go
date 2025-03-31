package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetAlbum(c *gin.Context, _ *utils.User, r struct {
	Path string `form:"path"`
}) {

	var album utils.Album
	if err := DB.
		Preload("Cover", utils.LoadFilename).
		Preload("Creator", utils.LoadOwnerInfo).
		Preload("Photos").
		Preload("Reviews").
		First(&album, "path = ?", r.Path).
		Error; err != nil {
		c.JSON(500, utils.Resp("获取相册失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取相册成功", nil, &album))
}
