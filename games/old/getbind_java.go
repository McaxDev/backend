package main

import (
	"context"
	"fmt"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetBindJava(c *gin.Context, user *dbs.User, name string) {

	if user.JavaName != nil {
		c.JSON(400, utils.Resp("你已经绑定了", nil, nil))
		return
	}

	authcode := utils.RandomCode(6, false)

	if err := Redis.Set(
		context.Background(), "auth_java_"+name, authcode, 10*time.Minute,
	).Err(); err != nil {
		c.JSON(500, utils.Resp("验证码存储失败", err, nil))
		return
	}

	_, err := Rcon.Execute(
		fmt.Sprintf("你的验证码是：%s，发送者：%s。", authcode, user.Name),
	)
	if err != nil {
		c.JSON(500, utils.Resp("消息发送失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("消息发送成功", nil, nil))
}
