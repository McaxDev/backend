package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	var data []dbs.Wiki
	if err := DB.Select(
		"ID", "Path", "Title", "Category",
	).Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("获取列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取列表成功", nil, data))
}
