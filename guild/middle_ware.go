package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/mids"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthGuild[T any](
	ajc mids.AuthJwtConfig,
	roles []uint,
	logicFunc func(c *gin.Context, user *dbs.User, req T),
) gin.HandlerFunc {
	return mids.AuthJwt(ajc, func(c *gin.Context, user *dbs.User, req T) {

		if !HandleAuthGuild(roles, user) {
			c.JSON(400, utils.Resp("你的公会权限不满足", nil, nil))
			return
		}

		logicFunc(c, user, req)
	})
}

func OnlyAuthGuild(
	ajc mids.AuthJwtConfig,
	roles []uint,
	logicFunc func(c *gin.Context, user *dbs.User),
) gin.HandlerFunc {
	return mids.OnlyAuthJwt(ajc, func(c *gin.Context, user *dbs.User) {

		if !HandleAuthGuild(roles, user) {
			c.JSON(400, utils.Resp("你的公会权限不满足", nil, nil))
			return
		}

		logicFunc(c, user)
	})
}

func HandleAuthGuild(roles []uint, user *dbs.User) bool {
	for _, value := range roles {
		if user.GuildRole == value {
			return true
		}
	}
	return false
}
