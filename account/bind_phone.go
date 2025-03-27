package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindPhone(c *gin.Context, u *utils.User, r struct {
	PhoneID string
}) {

	if err := DB.First(
		new(utils.User), "phone = ?", r.PhoneID,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此手机已被绑定", nil, nil))
		return
	}

	if err := DB.Model(&u).Update(
		"phone", r.PhoneID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新成功", nil, nil))
}
