package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JoinGuild(c *gin.Context, user *dbs.User, id uint) {

	if err := DB.First(
		new(dbs.Guild), "id = ?", id,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("不存在此公会", err, nil))
		return
	}

	user.GuildID = &id
	user.GuildRole = 1

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(500, utils.Resp("公会加入失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("公会加入成功", nil, nil))
}
