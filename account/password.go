package main

import (
	"context"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResetPassword(c *gin.Context) {

	var request struct {
		Codetype string
		Number   string
		Authcode string
		Password string
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	if _, err := AuthClient.Auth(
		context.Background(),
		&rpc.Authcode{
			Codetype: request.Codetype,
			Number:   request.Number,
			Authcode: request.Authcode,
		},
	); err != nil {
		c.JSON(400, utils.Resp("验证失败", err, nil))
		return
	}

	numberType := "telephone"
	if request.Codetype == "email" {
		numberType = "email"
	}

	if err := DB.Model(new(dbs.User)).Where(
		numberType+" = ?", request.Number,
	).Update(
		"Password", request.Password,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("不存在这个用户", nil, nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("密码修改失败", err, nil)) 
		return
	}

	c.JSON(200, utils.Resp("密码修改成功", nil, nil))
}
