package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context, req struct {
	Path string `form:"path"`
}) {

	var album utils.Album
	if err := DB.First(&album, "path = ?", req.Path).Error; err != nil {
		c.JSON(500, utils.Resp("查找相册失败", err, nil))
		return
	}

	var data []utils.Image
	if err := DB.Where("album_id = ?", album.ID).Find(
		&data,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取图片列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("", nil, data))
}
