package main

import (
	"fmt"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context, req struct {
	ID uint `form:"id"`
}) {

	var data []dbs.Image
	fmt.Println(req.ID)
	if err := DB.Where("album_id = ?", req.ID).Find(
		&data,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取图片列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取图片列表成功", nil, data))
}
