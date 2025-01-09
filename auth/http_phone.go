package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/gin-gonic/gin"
)

func SendPhone(c *gin.Context, user *dbs.User, req struct{}) {

	telephone := c.Param("number")
	authcode := utils.RandomCode(6, false)

	message := unisms.BuildMessage()
	message.SetTo(telephone)
	message.SetSignature(Config.SMS.Signature)
	message.SetTemplateId(Config.SMS.Template)
	message.SetTemplateData(map[string]string{
		"code": authcode,
		"ttl":  "10",
	})

	_, err := unisms.NewClient(
		Config.SMS.ID, Config.SMS.Secret,
	).Send(message)
	if err != nil {
		c.JSON(500, utils.Resp("短信发送失败", err, nil))
		return
	}

	if err := Redis.Set(
		context.Background(),
		"auth_phone_"+telephone, authcode, 10*time.Minute,
	).Err(); err != nil {
		c.JSON(500, utils.Resp("验证码存储失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("验证码发送成功", nil, nil))
}
