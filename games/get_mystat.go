package main

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyStat(c *gin.Context, u *dbs.User, r struct {
	Server string `form:"server"`
	IsJava bool   `form:"is_java"`
}) {

	var player string
	if r.IsJava && u.JavaName != nil {
		player = *u.JavaName
	} else if !r.IsJava && u.BedrockName != nil {
		player = *u.BedrockName
	} else {
		c.JSON(400, utils.Resp("你还未绑定玩家", nil, nil))
		return
	}

	uuid := Redis.HGet(
		context.Background(), r.Server+":name_uuid", player,
	)
	if err := uuid.Err(); err != nil {
		c.JSON(500, utils.Resp("查不到你的UUID", err, nil))
		return
	}

	file, err := os.Open(filepath.Join(
		Servers[r.Server].Path, "world/stats/", uuid.Val()+".json",
	))
	if err != nil {
		c.JSON(500, utils.Resp("查不到你的玩家数据", err, nil))
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"msg": "读取数据失败"})
		return
	}
	c.Data(200, "application/json", data)
}
