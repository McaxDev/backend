package main

import (
	"errors"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, req struct {
	Account  string
	Password string
	Authcode string
}) {

	var user dbs.User
	var accountType string
	if isPhone(req.Account) {
		user.Phone = &req.Account
		accountType = "phone"
	} else if isEmail(req.Account) {
		user.Email = req.Account
		accountType = "email"
	} else {
		user.Name = req.Account
	}

	if req.Authcode != "" {

		if accountType == "" {
			c.JSON(400, utils.Resp("请提供有效的号码", nil, nil))
			return
		}

		if err := Author.Auth(
			req.Account,
			req.Authcode,
			accountType,
		); err != nil {
			c.JSON(400, utils.Resp("验证码验证失败", err, nil))
			return
		}

		if err := DB.First(&user).Error; errors.Is(
			err, gorm.ErrRecordNotFound,
		) {
			if err := DB.Create(&user).Error; err != nil {
				c.JSON(500, utils.Resp("新用户创建失败", err, nil))
				return
			}
		} else if err != nil {
			c.JSON(500, utils.Resp("用户查询失败", err, nil))
			return
		}
	} else if req.Password != "" {

		if err := DB.First(&user).Error; errors.Is(
			err, gorm.ErrRecordNotFound,
		) {
			c.JSON(400, utils.Resp("你尚未注册", nil, nil))
			return
		} else if err != nil {
			c.JSON(500, utils.Resp("用户查询失败", err, nil))
			return
		}

		if req.Password != user.Password {
			c.JSON(400, utils.Resp("密码不正确", nil, nil))
			return
		}
	} else {
		c.JSON(400, utils.Resp("请至少提供一种验证方式", nil, nil))
		return
	}

	token, err := utils.GetJwt(user.ID, Config.JWTKey)
	if err != nil {
		c.JSON(500, utils.Resp("用户凭证生成失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("登录成功", nil, token))
}
