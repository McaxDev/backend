package main

import (
	"fmt"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendQQMailCode(c *gin.Context) {

	qq := c.Param("number")
	authcode := utils.RandomCode(6)
	expiry := time.Now().Add(10 * time.Minute)

	if err := SendEmailCode(
		fmt.Sprintf("%s@qq.com", qq), authcode, c.ClientIP(), expiry,
	); err != nil {
		c.JSON(500, utils.Resp("验证码发送失败", err, nil))
		return
	}

	QQMailSent.lock.Lock()
	QQMailSent.data[qq] = MsgSentValue{
		Authcode: authcode,
		Expiry:   expiry,
	}
	QQMailSent.lock.Unlock()

	c.JSON(200, utils.Resp("验证码发送成功", nil, nil))
}
