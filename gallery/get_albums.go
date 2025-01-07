package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {

	var data []dbs.Album
	if err := DB.Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("获取相册列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取相册列表成功", nil, nil))
}
