package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendQQCode(c *gin.Context) {

	qq := c.Param("number")
	authcode := utils.RandomCode(6, true)

	if err := Redis.Set(
		context.Background(),
		"auth_qq_"+qq, authcode, 10*time.Second,
	).Err(); err != nil {
		c.JSON(500, utils.Resp("验证码存储失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("请求验证码成功", nil, authcode))
}
