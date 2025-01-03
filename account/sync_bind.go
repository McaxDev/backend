package main

import (
	"fmt"
	"strings"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SyncBind(user *User, c *gin.Context) {

	isQQMail := strings.HasSuffix(user.Email, "@qq.com")
	qqMail := fmt.Sprintf("%s@qq.com", user.QQ)
	if user.QQ != "" && user.Email != "" {
		if isQQMail {
			c.JSON(400, utils.Resp("你已经绑定过了", nil, nil))
			return
		} else {
			user.Email = qqMail
		}
	} else if user.QQ == "" && user.Email != "" {
		if isQQMail {
			user.QQ = strings.Split(user.Email, "@")[0]
		} else {
			c.JSON(400, utils.Resp("你的邮箱不是QQ邮箱", nil, nil))
			return
		}
	} else if user.QQ != "" && user.Email == "" {
		user.Email = qqMail
	} else {
		c.JSON(400, utils.Resp("你没有绑定任何账号", nil, nil))
		return
	}
	if err := DB.Save(&user).Error; err != nil {
		c.JSON(500, utils.Resp("数据库修改失败", err, nil))
		return
	}
	c.JSON(200, utils.Resp("同步成功", nil, nil))
}
