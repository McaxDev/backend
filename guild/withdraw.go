package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Withdraw(user *dbs.User, c *gin.Context) {

	var amount uint
	if err := utils.GetBodyByCtx(c, &amount); err != nil {
		c.JSON(400, utils.Resp("无法读取请求", err, nil))
		return
	}

	if user.Guild.Money < amount {
		c.JSON(400, utils.Resp("公会余额不足", nil, nil))
		return
	}
}
