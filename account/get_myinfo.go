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
	Type  string
	Value string
}) {

	query := DB.Model(&user)

	switch req.Type {
	case "avatar":
		query = query.Update("Avatar", req.Value)
	case "profile":
		query = query.Update("Profile", req.Value)
	default:
		c.JSON(400, utils.Resp("不支持设置此种信息", nil, nil))
		return
	}

	if err := query.Error; err != nil {
		c.JSON(500, utils.Resp("资料更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("资料更新成功", nil, nil))
}
