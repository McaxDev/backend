package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyGuild(c *gin.Context, user *dbs.User) {
	c.JSON(200, utils.Resp("获取成功", nil, &user.Guild))
}
