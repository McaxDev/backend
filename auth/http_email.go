package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthEmail(c *gin.Context) {

	email := c.Param("number")
	authcode := utils.RandomCode(6, true)
	expiry := time.Now().Add(10 * time.Minute)

	if err := SendEmailCode(
		email, authcode, c.ClientIP(), expiry,
	); err != nil {
		c.JSON(500, utils.Resp("邮件发送失败", err, nil))
		return
	}

	if err := Redis.Set(
		context.Background(),
		"auth_email_"+email, authcode, 10*time.Minute,
	).Err(); err != nil {
		c.JSON(500, utils.Resp("验证码存储失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("邮件发送成功", nil, nil))
}
