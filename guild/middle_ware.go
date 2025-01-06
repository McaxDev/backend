package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthJoin[T any](join bool, logicFunc utils.LogicFunc[T]) gin.HandlerFunc {
	return utils.AuthJwt(func(user *dbs.User, c *gin.Context, req T) {

		if user.GuildID == 0 && join {
			c.JSON(400, utils.Resp("你还没有加入公会", nil, nil))
			return
		}

		if user.GuildID != 0 && !join {
			c.JSON(400, utils.Resp("你已经加入公会了", nil, nil))
			return
		}

		logicFunc(user, c, req)
	})
}

func AuthRole[T any](role uint, logicFunc utils.LogicFunc[T]) gin.HandlerFunc {
	return AuthJoin(true, func(user *dbs.User, c *gin.Context, req T) {

		if user.GuildRole < role {
			c.JSON(403, utils.Resp("你不是公会管理员", nil, nil))
			return
		}

		logicFunc(user, c, req)
	})
}
