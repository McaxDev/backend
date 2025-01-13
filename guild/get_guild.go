package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetGuild(c *gin.Context, guildId uint) {

	var guild dbs.Guild
	if err := DB.Preload("Users").Where(
		"ID = ?", guildId,
	).First(&guild).Error; err != nil {
		c.JSON(500, utils.Resp("查看公会失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查看公会成功", nil, &guild))
}
