package main

import (
	"context"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthQRCode(c *gin.Context, u *utils.User, r string) {

	result, err := Redis.Get(
		context.Background(), "auth_qr_"+r,
	).Result()
	if err != nil || result != "true" {
		c.JSON(400, utils.Resp("二维码无效", err, nil))
		return
	}

}
