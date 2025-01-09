package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Delete(user *dbs.User, c *gin.Context, req uint) {

	if err := DB.Where("id = ?", req).Delete(
		new(dbs.Wiki),
	).Error; err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("数据删除成功", nil, nil))
}
