package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

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

	user.ClearPrivate()

	c.JSON(200, utils.Resp("查询成功", nil, user))
}
