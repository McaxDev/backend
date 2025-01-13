package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Appoint(c *gin.Context, user *dbs.User, req struct {
	IDs   []uint
	Agree bool
}) {

	query := DB.Model(new(dbs.User)).Where(
		"guild_id = ? AND id IN ?", user.GuildID, req.IDs,
	)
	var err error

	if req.Agree {
		err = query.Where("guild_role = ?", 2).Update("guild_role", 3).Error
	} else {
		err = query.Where("guild_role = ?", 3).Update("guild_role", 2).Error
	}
	if err != nil {
		c.JSON(500, utils.Resp("操作失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("任命管理员成功", nil, nil))
}
