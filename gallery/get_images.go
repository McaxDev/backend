package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context, ID uint) {

	var data dbs.Image
	if err := DB.Where("id = ?", ID).First(
		&data,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取图片列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取图片列表成功", nil, &data))
}
