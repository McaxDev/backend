package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindEmail(c *gin.Context, u *utils.User, r struct {
	EmailID string
}) {

	if err := DB.First(
		new(utils.User), "email = ?", r.EmailID,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此邮箱已被绑定", err, nil))
		return
	}

	if err := DB.Model(&u).Update(
		"email", r.EmailID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新成功", nil, nil))
}
