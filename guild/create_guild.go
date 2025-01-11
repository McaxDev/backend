package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateGuild(user *dbs.User, c *gin.Context, req struct {
	GID  string
	Name string
}) {

	if err := DB.Where(
		"gid = ?", req.GID,
	).Or(
		"name = ?", req.Name,
	).First(new(dbs.Guild)).Error; err != nil {
		c.JSON(400, utils.Resp("此公会已存在", err, nil))
		return
	}

	if err := user.ExecWithCoins(
		DB, 5, false, func(tx *gorm.DB) error {
			return tx.Create(&dbs.Guild{
				GID:    req.GID,
				Name:   req.Name,
				Number: 1,
			}).Error
		},
	); err != nil {
		c.JSON(400, utils.Resp("公会创建失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("公会创建成功", nil, nil))
}
