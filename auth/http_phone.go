package main

import (
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/limiter"
	"github.com/McaxDev/backend/utils"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/gin-gonic/gin"
)

func SendPhone(user *dbs.User, c *gin.Context) {

	telephone := c.Param("number")
	authcode := utils.RandomCode(6, false)

	if err := limiter.UseLimiter(user.Name, "phone"); err != nil {
		c.JSON(400, utils.Resp("请求频率太快", err, nil))
		return
	}

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

	PhoneSent.lock.Lock()
	PhoneSent.data[telephone] = MsgSentValue{
		Authcode: authcode,
		Expiry:   time.Now().Add(10 * time.Minute),
	}
	PhoneSent.lock.Unlock()

	c.JSON(200, utils.Resp("验证码发送成功", nil, nil))
}
