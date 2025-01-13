package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Leave(c *gin.Context, user *dbs.User) {

	user.GuildID = nil
	user.GuildRole = 0
	if err := DB.Updates(user).Error; err != nil {
		c.JSON(500, utils.Resp("公会退出失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("操作成功", nil, nil))
}
