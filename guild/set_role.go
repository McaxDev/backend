package main

import (
	"errors"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRole(user *dbs.User, c *gin.Context, req struct {
	UserIDs []uint
	Role    uint
}) {

	if req.Role >= user.GuildRole {
		c.JSON(403, utils.Resp("你的权限不足", nil, nil))
		return
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {

		tx = tx.Model(new(dbs.User)).Where(
			"guild_id = ?", user.GuildID,
		)

		if err := tx.Where(
			"id IN ?", req.UserIDs,
		).Update(
			"guild_role = ?", req.Role,
		).Error; err != nil {
			return err
		}

		if err := tx.Where("guild_role = 3").First(
			new(dbs.User),
		).Error; err == gorm.ErrRecordNotFound {
			return errors.New("至少要有一个公会长")
		}

		return nil
	}); err != nil {
		c.JSON(500, utils.Resp("修改权限失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("修改权限成功", nil, nil))
}
