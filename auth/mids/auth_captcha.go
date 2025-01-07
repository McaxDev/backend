package mids

import (
	"github.com/McaxDev/backend/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func AuthCaptcha(c *gin.Context) {

	var req struct {
		CaptchaID    string
		CaptchaValue string
	}
	if err := utils.GetBodyByCtx(c, &req); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	if !captcha.VerifyString(req.CaptchaID, req.CaptchaValue) {
		c.JSON(400, utils.Resp("验证码不正确", nil, nil))
		return
	}

	c.Next()
}
