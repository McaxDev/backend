package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyinfo(c *gin.Context, user *dbs.User) {

	c.JSON(200, utils.Resp("获取成功", nil, user))
}

func SetUserInfo(c *gin.Context, user *dbs.User, req struct {
	Avatar  string
	Profile string
}) {

	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if req.Profile != "" {
		user.Profile = req.Profile
	}

	if err := DB.Save(user).Error; err != nil {
		c.JSON(500, utils.Resp("资料更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("资料更新成功", nil, nil))
}
