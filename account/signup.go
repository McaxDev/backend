package main

import (
	"context"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context, req struct {
	Username  string
	Password  string
	Email     string
	EmailCode string
}) {

	if err := DB.First(
		new(dbs.User), "username = ?", req.Username,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此用户已经注册过了", nil, nil))
		return
	}

	if err := DB.First(
		new(dbs.User), "email = ?", req.Email,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此邮箱已经注册过了", nil, nil))
		return
	}

	if _, err := AuthClient.Auth(
		context.Background(),
		&rpc.Authcode{
			Codetype: "email",
			Number:   req.Email,
			Authcode: req.EmailCode,
		},
	); err != nil {
		c.JSON(400, utils.Resp("邮箱验证失败", err, nil))
		return
	}

	user := dbs.User{
		Name:     req.Username,
		Password: req.Password,
		Email:    req.Email,
		Admin:    false,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, utils.Resp("无法在数据库里创建新用户", err, nil))
		return
	}

	token, err := utils.GetJwt(user.ID)
	if err != nil {
		c.JSON(500, utils.Resp("JWT生成失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("用户创建成功", nil, gin.H{
		"token": token,
	}))
}
