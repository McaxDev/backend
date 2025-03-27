package main

import (
	"fmt"
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMyinfo(c *gin.Context, u *utils.User, r struct{}) {

	fmt.Println(u.TempMeta)
	c.JSON(200, utils.Resp("获取成功", nil, u))
}

func GetUserinfo(c *gin.Context, _ *utils.User, r struct {
	ID uint `form:"id"`
}) {

	var user utils.User
	if err := DB.Preload("Guild").Preload(
		"Props",
	).Preload("Comments").Preload("Albums").First(
		&user, "id = ?", r.ID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("查询用户失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查询成功", nil, user))
}

func SetUserInfo(c *gin.Context, u *utils.User, r struct {
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
	Profile  string    `json:"profile"`
}) {

	if r.Avatar != "" {
		u.Avatar = r.Avatar
	}

	if r.Profile != "" {
		u.Profile.Content = r.Profile
	}

	if err := DB.Save(u).Error; err != nil {
		c.JSON(500, utils.Resp("资料更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("资料更新成功", nil, nil))
}
