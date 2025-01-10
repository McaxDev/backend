package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindPhone(c *gin.Context, user *dbs.User, req struct {
	PhoneID string
}) {

	if err := DB.First(
		new(dbs.User), "phone = ?", req.PhoneID,
	).Error; err == nil {
		c.JSON(400, utils.Resp("此手机已被绑定", nil, nil))
		return
	}

	if err := DB.Model(&user).Update(
		"phone", req.PhoneID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新成功", nil, nil))
}
