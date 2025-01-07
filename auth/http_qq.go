package main

import (
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SendQQCode(c *gin.Context) {

	qq := c.Param("number")
	authcode := utils.RandomCode(6, true)
	expiry := time.Now().Add(10 * time.Minute)

	QQSent.lock.Lock()
	QQSent.data[qq] = MsgSentValue{
		Authcode: authcode,
		Expiry:   expiry,
	}
	QQSent.lock.Unlock()

	c.JSON(200, utils.Resp("请求验证码成功", nil, authcode))
}
