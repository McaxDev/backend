package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Donate(c *gin.Context, user *dbs.User, amount uint) {

	if err := user.ExecWithCoins(
		DB, amount, true, func(tx *gorm.DB) error {
			return tx.Model(new(dbs.Guild)).Where(
				"id = ?", user.GuildID,
			).Update(
				"money = ?", user.Guild.Money+amount,
			).Error
		},
	); err != nil {
		c.JSON(500, utils.Resp("捐款失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("捐款成功", nil, nil))
}
