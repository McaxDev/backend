package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Signout(c *gin.Context, u *utils.User, _ struct{}) {

	if err := DB.Delete(&u).Error; err != nil {
		c.JSON(500, utils.Resp("注销失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("注销成功，感谢使用", nil, nil))
}
