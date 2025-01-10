package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/McaxDev/backend/utils/auth"
	"github.com/gin-gonic/gin"
)

func Signout(c *gin.Context, user *dbs.User, req auth.CodeMsg) {

	if err := DB.Delete(&user).Error; err != nil {
		c.JSON(500, utils.Resp("注销失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("注销成功，感谢使用", nil, nil))
}
