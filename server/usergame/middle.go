package main

import (
	"context"

	"github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

type LogicFunc func(user *rpc.Userinfo, c *gin.Context)

func AuthJwtMiddle(logicFunc LogicFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		user, err := AccountClient.VerifyJwt(
			context.Background(),
			&rpc.JwtRequest{JWT: c.GetHeader("Authorization")},
		)
		if err != nil {
			c.JSON(500, utils.Resp("身份验证服务器错误", err))
			return
		}

		logicFunc(user, c)
	}
}

func AuthAdminMiddle(logicFunc LogicFunc) LogicFunc {
	return func(user *rpc.Userinfo, c *gin.Context) {

		if !user.Admin {
			c.JSON(400, utils.Resp("你不是管理员", nil))
			return
		}

		logicFunc(user, c)
	}
}
