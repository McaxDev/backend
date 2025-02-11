package main

import (
	"context"

	account "github.com/McaxDev/backend/account/rpc"
	gameapi "github.com/McaxDev/backend/gameapi/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendCmd(user *account.Userinfo, c *gin.Context) {

	var request struct {
		Server  string
		Command string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	response, err := GameapiClient.SendCmd(
		context.Background(), &gameapi.CmdReq{
			Server: request.Server,
			Cmd:    request.Command,
		},
	)
	if err != nil {
		c.JSON(500, utils.Resp("命令发送失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("命令执行成功", nil, response.Data))
}
