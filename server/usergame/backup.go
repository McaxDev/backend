package main

import (
	"context"

	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func WorldBackup(user *account.User, c *gin.Context) {

	var request struct {
		Server string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("读取请求参数失败", err, nil))
		return
	}

	response, err := GameapiClient.WorldBackup(
		context.Background(),
		&gameapirpc.BackupRequest{Server: request.Server},
	)
	if err != nil || !response.Success {
		c.JSON(500, utils.Resp("备份失败", err))
		return
	}

	c.JSON(200, utils.Resp("备份成功", nil))
}
