package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func SetWiki(user *dbs.User, c *gin.Context, req dbs.Wiki) {

	if err := DB.Where(
		"path = ?", req.Path,
	).Or(
		"title = ?", req.Title,
	).Or(
		"order = ?", req.Order,
	).Find(new(dbs.Wiki)).Error; err == nil {
		c.JSON(400, utils.Resp("此目录已存在", nil, nil)) 
		return
	}

	req.HTML = string(blackfriday.Run([]byte(req.Markdown)))

	if err := DB.Save(&req).Error; err != nil {
		c.JSON(500, utils.Resp("设置文档失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置文档成功", nil, nil))
}
