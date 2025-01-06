package mids

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
)

func AuthAdmin[T any](
	logicFunc LogicFunc[T], preloads ...string,
) gin.HandlerFunc {
	return AuthJwt(
		func(user *dbs.User, c *gin.Context, params T) {

			if !user.Admin {
				c.JSON(403, Resp("你不是管理员", nil, nil))
				return
			}

			logicFunc(user, c, params)
		},
		preloads...,
	)
}
