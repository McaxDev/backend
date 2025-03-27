package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func UnbindEmail(c *gin.Context, user *utils.User, req struct {
	EmailID string
}) {

	if req.EmailID != user.Email {
		c.JSON(400, utils.Resp("这不是你的邮箱", nil, nil))
		return
	}

	if err := DB.Model(user).Update(
		"email", "",
	).Error; err != nil {
		c.JSON(500, utils.Resp("邮箱更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("邮箱更新成功", nil, nil))
}
