package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context, user *utils.User, req uint) {

	if err := DB.Where("id = ?", req).Delete(
		new(utils.Wiki),
	).Error; err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("数据删除成功", nil, nil))
}
