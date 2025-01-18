package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetBindBedrock(c *gin.Context, user *dbs.User, req struct {
	Username  string
	AccessKey string
}) {

	if Config.AccessKey != req.AccessKey {
		c.String(400, "非法访问")
		return
	}

	authcode := utils.RandomCode(6, false)

	if err := Redis.Set(
		context.Background(),
		"auth_bedrock_"+req.Username,
		authcode,
		10*time.Minute,
	).Err(); err != nil {
		c.String(500, "创建验证码失败")
		return
	}

	c.String(200, "验证码创建成功："+authcode)
}
