package main

import (
	"fmt"

	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/database"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Edit(user *account.User, c *gin.Context) {

	table := c.Query("table")
	var err error
	switch table {
	case "category":
		err = HandleEdit[database.Category](c)
	case "wiki":
		err = HandleEdit[database.Wiki](c)
	default:
		err = fmt.Errorf("无法修改：%s\n", table)
	}
	if err != nil {
		c.JSON(400, utils.Resp("修改失败", err, nil))
		return
	}
	c.JSON(200, utils.Resp("修改成功", nil, nil))
}

func HandleEdit[T any](c *gin.Context) error {

	var req database.Category
	if err := utils.GetBodyByCtx(c, &req); err != nil {
		return err
	}

	if err := DB.Where(
		"path = ?", req.Path,
	).Or(
		"title = ?", req.Title,
	).Or(
		"order = ?", req.Order,
	).Find(new(database.Category)).Error; err == nil {
		return err
	}

	return DB.Save(&req).Error
}
