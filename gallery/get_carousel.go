package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetCarousel(c *gin.Context) {
	var images []dbs.Image
	if err := DB.Find(&images, "album_id = ?", 2).Error; err != nil {
		c.JSON(500, utils.Resp("获取失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("", nil, images))
}
