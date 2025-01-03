package utils

import (
	"context"

	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

type LogicFunc func(user *account.User, c *gin.Context)

func AuthJwt(logicFunc LogicFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if len(token) < 8 {
			ctx.JSON(401, utils.Resp("token无效", nil, nil))
			return
		}

		resp, err := AccountClient.GetUser(
			context.Background(),
			&account.JWT{JWT: token[7:]},
		)
		if err != nil {
			ctx.JSON(401, utils.Resp("验证失败", err, nil))
			return
		}

		logicFunc(resp, ctx)
	}
}

func AuthAdmin(logicFunc LogicFunc) gin.HandlerFunc {
	return AuthJwt(func(user *account.User, c *gin.Context) {

		if !user.Admin {
			c.JSON(403, utils.Resp("你不是管理员", nil, nil))
			return
		}

		logicFunc(user, c)
	})
}
