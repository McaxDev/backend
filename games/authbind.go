package main

import (
	"context"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthBind(c *gin.Context, user *dbs.User, r struct {
	Player   string `json:"player"`
	Game     string `json:"game"`
	Authcode string `json:"authcode"`
}) {

	authcode, err := Redis.Get(
		context.Background(), "auth_"+r.Game+"_"+r.Player,
	).Result()
	if err != nil {
		c.JSON(400, utils.Resp("没有验证码", err, nil))
		return
	}

	if authcode != r.Authcode {
		c.JSON(400, utils.Resp("验证码不正确", nil, nil))
		return
	}

	if err := DB.Model(user).Update(
		r.Game+"_name", r.Player,
	).Error; err != nil {
		c.JSON(500, utils.Resp("绑定失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("绑定成功", nil, nil))
}
