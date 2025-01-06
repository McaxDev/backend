package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditGuildName(user *dbs.User, c *gin.Context) {

	var req struct {
		GID  string
		Name string
	}
	if err := utils.GetBodyByCtx(c, &req); err != nil {
		c.JSON(400, utils.Resp("用户请求有误", err, nil))
		return
	}

	if err := DB.Where(
		"gid = ?", req.GID,
	).Or(
		"name = ?", req.Name,
	).First(new(dbs.Guild)).Error; err == nil {
		c.JSON(400, utils.Resp("此公会已存在", nil, nil))
		return
	}

	if err := utils.ExecWithCoins(user, 1, func(tx *gorm.DB) error {
		return tx.Where("id = ?", user.Guild.ID).Updates(&dbs.Guild{
			GID:  req.GID,
			Name: req.Name,
		}).Error
	}); err != nil {
		c.JSON(400, utils.Resp("公会名修改失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("公会名修改成功", nil, nil))
}
