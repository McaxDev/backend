package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transfer(c *gin.Context, user *dbs.User, id uint) {

	if err := DB.First(
		new(dbs.User), "id = ? AND guild_id = ?", id, user.GuildID,
	).Error; err != nil {
		c.JSON(400, utils.Resp("找不到这个用户", err, nil))
		return
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {

		tx = tx.Model(new(dbs.User))
		if err := tx.Where("id = ?", user.ID).Update(
			"guild_role", 3,
		).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", id).Update(
			"guild_role", 4,
		).Error
	}); err != nil {
		c.JSON(500, utils.Resp("移交权限失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("移交权限成功", nil, nil))
}
