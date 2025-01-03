package main

import (
	"github.com/McaxDev/backend/database"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

	var resp []database.Category
	if err := DB.Preload("Wiki").Find(&resp).Error; err != nil {
		c.JSON(500, utils.Resp("获取wiki失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取wiki成功", nil, resp))
}
