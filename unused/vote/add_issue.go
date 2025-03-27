package main

import (
	"context"

	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AddIssue(user *account.User, c *gin.Context) {

	var request struct {
		Title   string
		Content string
	}
	if err := utils.LoadBodyByCtx(c, &request); err != nil {
		c.JSON(400, utils.Resp("读取请求失败", err, nil))
		return
	}

	if _, err := AccountClient.CostMoney(
		context.Background(), &account.Int{Data: 1},
	); err != nil {
		c.JSON(400, utils.Resp("扣款失败", err, nil))
		return
	}

}
