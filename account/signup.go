package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context, _ *utils.User, r struct {
	Username string
	Password string
	EmailID  string
}) {

	if err := DB.First(
		new(utils.User), "username = ?", r.Username,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此用户已经注册过了", nil, nil))
		return
	}

	if err := DB.First(
		new(utils.User), "email = ?", r.EmailID,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此邮箱已经注册过了", nil, nil))
		return
	}

	user := utils.User{
		Name:     r.Username,
		Password: r.Password,
		Email:    r.EmailID,
		Admin:    false,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, utils.Resp("无法在数据库里创建新用户", err, nil))
		return
	}

	token, err := utils.GetJwt(user.ID, Config.JWTKey)
	if err != nil {
		c.JSON(500, utils.Resp("JWT生成失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("用户创建成功", nil, token))
}
