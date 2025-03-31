package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetDocument(c *gin.Context, _ *utils.User, r struct {
	Path string `form:"path"`
}) {

	var data utils.Document
	if err := DB.First(&data, "path = ?", r.Path).Error; err != nil {
		c.JSON(500, utils.Resp("查询失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查询成功", nil, data))
}
