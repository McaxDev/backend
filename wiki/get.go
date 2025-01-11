package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context, req struct {
	ID uint `form:"id"`
}) {

	var data dbs.Wiki
	if err := DB.Where("id = ?", req.ID).First(
		&data,
	).Error; err != nil {
		c.JSON(500, utils.Resp("wiki获取失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("wiki获取成功", nil, &data))
}
