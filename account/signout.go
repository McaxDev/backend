package main

import (
	"context"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Signout(user *dbs.User, c *gin.Context, req *rpc.Authcode) {

	if _, err := AuthClient.Auth(
		context.Background(), req,
	); err != nil {
		c.JSON(400, utils.Resp("联系方式验证失败", err, nil))
		return
	}

	if err := DB.Delete(&user).Error; err != nil {
		c.JSON(500, utils.Resp("注销删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("注销成功，感谢使用", nil, nil))
}
