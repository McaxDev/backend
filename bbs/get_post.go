package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context, r struct {
	ID uint `form:"id"`
}) {

	var post dbs.Post
	if err := DB.Preload("User").Preload("Comments").First(
		&post, "id = ?", r.ID,
	).Error; err != nil {
		c.JSON(500, utils.Resp("查询帖子失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取成功", nil, &post))
}
