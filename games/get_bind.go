package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetBind(c *gin.Context, r struct {
	Player    string
	Game      string
	AccessKey string
}) {

	if Config.AccessKey != r.AccessKey {
		c.String(400, "非法访问")
		return
	}

	authcode := utils.RandomCode(6, false)

	if err := Redis.Set(
		context.Background(),
		"auth_"+r.Game+"_"+r.Player,
		authcode,
		10*time.Minute,
	).Err(); err != nil {
		c.String(500, "创建验证码失败")
		return
	}

	c.String(200, "验证码创建成功："+authcode)
}
