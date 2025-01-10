package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upgrade(user *dbs.User, c *gin.Context) {

	var cost uint
	switch user.Guild.Level {
	case 1:
		cost = 100
	case 2:
		cost = 200
	case 3:
		c.JSON(400, utils.Resp("你的公会已经满级了", nil, nil))
		return
	default:
		c.JSON(500, utils.Resp("公会等级异常", nil, nil))
		return
	}

	if err := user.ExecWithCoins(
		DB, cost, false, func(tx *gorm.DB) error {
			return tx.Model(&user.Guild).Update(
				"level = ?", user.Guild.Level+1,
			).Error
		},
	); err != nil {
		c.JSON(400, utils.Resp("公会升级失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("公会升级成功", nil, nil))
}
