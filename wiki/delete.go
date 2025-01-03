package main

import (
	"fmt"

	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/database"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Delete(user *account.User, c *gin.Context) {

	var req uint
	if err := utils.GetBodyByCtx(c, &req); err != nil {
		c.JSON(400, utils.Resp("请求格式有误", err, nil))
		return
	}

	var err error
	switch table := c.Param("table"); table {
	case "category":
		err = HandleDelete[database.Category](req)
	case "wiki":
		err = HandleDelete[database.Wiki](req)
	default:
		err = fmt.Errorf("不支持删除：%s\n", table)
	}
	if err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("数据删除成功", nil, nil))
}

func HandleDelete[T any](id uint) error {
	return DB.Where("id = ?", id).Delete(new(T)).Error
}
