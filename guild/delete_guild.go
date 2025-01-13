package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Dissolve(c *gin.Context, user *dbs.User) {

	if err := DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where(
			"guild_id = ?", user.GuildID,
		).Updates(&dbs.User{
			GuildID:   nil,
			GuildRole: 0,
		}).Error; err != nil {
			return err
		}

		return tx.Delete(&user.Guild).Error

	}); err != nil {
		c.JSON(500, utils.Resp("公会解散失败", err, nil))
	}

	c.JSON(200, utils.Resp("操作成功", nil, nil))
}
