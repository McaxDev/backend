package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func UnbindPhone(c *gin.Context, user *utils.User, req struct {
	PhoneID string
}) {

	if req.PhoneID != user.Email {
		c.JSON(400, utils.Resp("这不是你的手机号", nil, nil))
		return
	}

	if err := DB.Model(user).Update(
		"phone", "",
	).Error; err != nil {
		c.JSON(500, utils.Resp("手机号更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("手机号更新成功", nil, nil))
}
