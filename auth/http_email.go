package main

import (
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthEmail(c *gin.Context) {

	email := c.Param("number")
	authcode := utils.RandomCode(6)
	expiry := time.Now().Add(10 * time.Minute)

	if err := SendEmailCode(
		email, authcode, c.ClientIP(), expiry,
	); err != nil {
		c.JSON(500, utils.Resp("邮件发送失败", err, nil))
		return
	}

	EmailSent.lock.Lock()
	EmailSent.data[email] = MsgSentValue{
		Authcode: authcode,
		Expiry:   expiry,
	}
	EmailSent.lock.Unlock()

	c.JSON(200, utils.Resp("邮件发送成功", nil, nil))
}
