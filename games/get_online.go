package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetOnline(c *gin.Context, r struct {
	Server string `form:"server"`
	Limit  uint   `form:"limit"`
}) {

	var result []dbs.Online
	if err := DB.Where("server = ?", r.Server).Order(
		"time DESC",
	).Limit(int(r.Limit)).Find(&result).Error; err != nil {
		c.JSON(500, utils.Resp("查找失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查找成功", nil, result))
}
