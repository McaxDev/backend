package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context, r struct {
	Category uint `form:"category"`
}) {

	var data []dbs.Post
	if err := DB.Preload("User").Preload("Comments").Select(
		"ID", "CreatedAt", "UpdatedAt", "Title",
	).Find(
		&data, "category = ?", r.Category,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取帖子列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取成功", nil, data))
}
