package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetThreads(c *gin.Context, u *dbs.User) {
	c.JSON(200, utils.Resp("获取成功", nil, u.Threads))
}
