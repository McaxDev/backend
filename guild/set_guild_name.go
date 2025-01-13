package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetGuildName(c *gin.Context, user *dbs.User, name string) {

	if err := DB.Where(
		"name = ?", name,
	).First(new(dbs.Guild)).Error; err == nil {
		c.JSON(400, utils.Resp("此公会已存在", nil, nil))
		return
	}

	if err := user.ExecWithCoins(
		DB, 1, false, func(tx *gorm.DB) error {
			return tx.Where("id = ?", user.Guild.ID).Updates(&dbs.Guild{
				Name: name,
			}).Error
		}); err != nil {
		c.JSON(400, utils.Resp("公会名修改失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("公会名修改成功", nil, nil))
}
