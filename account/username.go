package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUsername(c *gin.Context, user *dbs.User, name string) {

	if err := c.BindJSON(&name); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	if err := DB.First(
		new(dbs.User), "username = ?", name,
	).Error; err == nil {
		c.JSON(400, utils.Resp("已经有人使用过这个名称了", nil, nil))
		return
	}

	if err := user.ExecWithCoins(DB, 100, false,
		func(tx *gorm.DB) error {
			return tx.Model(&user).Update("name", name).Error
		},
	); err != nil {
		c.JSON(500, utils.Resp("用户名更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("用户名更新成功", nil, nil))
}
