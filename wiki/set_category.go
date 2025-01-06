package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetCategory(user *dbs.User, c *gin.Context, req dbs.Category) {

	if err := DB.Where(
		"path = ?", req.Path,
	).Or(
		"title = ?", req.Title,
	).Or(
		"order = ?", req.Order,
	).Find(new(dbs.Category)).Error; err == nil {
		c.JSON(400, utils.Resp("此分类已存在", nil, nil)) 
		return
	}

	if err := DB.Save(&req).Error; err != nil {
		c.JSON(500, utils.Resp("设置分类失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置分类成功", nil, nil))
}
