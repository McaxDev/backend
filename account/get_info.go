package main

import (
	"fmt"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyinfo(c *gin.Context, user *dbs.User) {

	fmt.Println(user.TempMeta)
	c.JSON(200, utils.Resp("获取成功", nil, user))
}

func GetUserinfo(c *gin.Context, req struct {
	ID uint `form:"id"`
}) {

	var user dbs.User
	if err := DB.Preload("Guild").Preload(
		"Props",
	).Preload("Comments").Preload("Albums").First(
		&user, "id = ?", req.ID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("查询用户失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查询成功", nil, user))
}

func SetUserInfo(c *gin.Context, user *dbs.User, req struct {
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
	Profile  string    `json:"profile"`
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
