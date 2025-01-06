package main

import (
	"fmt"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Delete(user *dbs.User, c *gin.Context, req uint) {

	var err error
	query := DB.Where("id = ?", req)
	switch table := c.Param("table"); table {
	case "category":
		err = query.Delete(new(dbs.Category)).Error
	case "wiki":
		err = query.Delete(new(dbs.Wiki)).Error
	default:
		err = fmt.Errorf("不支持删除：%s\n", table)
	}
	if err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("数据删除成功", nil, nil))
}
