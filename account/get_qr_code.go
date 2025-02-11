package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetQRCode(c *gin.Context) {
	seed := utils.RandomCode(6, true)
	if err := Redis.Set(
		context.Background(),
		"auth_qr_"+seed, true, 10*time.Minute,
	).Err(); err != nil {
		c.JSON(500, utils.Resp("存储二维码失败", err, nil))
		return
	}
	c.JSON(200, utils.Resp("", nil, seed))
}
