package main

import (
	"context"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

type AuthBindReq struct {
	Gamename string
	Authcode string
}

func AuthBind(
	game string,
) func(c *gin.Context, user *dbs.User, req AuthBindReq) {
	return func(c *gin.Context, user *dbs.User, req AuthBindReq) {

		authcode, err := Redis.Get(
			context.Background(), "auth_"+game+"_"+req.Gamename,
		).Result()
		if err != nil {
			c.JSON(400, utils.Resp("没有验证码", err, nil))
			return
		}

		if authcode != req.Authcode {
			c.JSON(400, utils.Resp("验证码不正确", nil, nil))
			return
		}

		if err := DB.Model(user).Update(
			game+"_name", req.Gamename,
		).Error; err != nil {
			c.JSON(500, utils.Resp("绑定失败", err, nil))
			return
		}

		c.JSON(200, utils.Resp("绑定成功", nil, nil))
	}
}
