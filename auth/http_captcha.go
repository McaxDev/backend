package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func SendCaptcha(c *gin.Context) {

	id := captcha.New()
	c.Header("Content-Type", "image/png")
	c.Header("X-Captcha-Id", id)
	if err := captcha.WriteImage(
		c.Writer, id, captcha.StdWidth, captcha.StdHeight,
	); err != nil {
		c.JSON(500, utils.Resp("验证码绘制失败", nil, nil))
		return
	}
}
