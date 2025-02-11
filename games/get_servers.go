package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetServers(c *gin.Context) {
	c.JSON(200, utils.Resp("", nil, Servers))
}
