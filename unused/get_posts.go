package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPosts(c *gin.Context, r struct {
	Category uint `form:"category"`
}) {

	var data []utils.Post
	if err := DB.Preload("User").Preload("Comments").Select(
		"ID", "CreatedAt", "UpdatedAt", "Title",
	).Find(
		&data, "category = ?", r.Category,
	).Error; err != nil {
		c.JSON(500, utils.Resp("获取帖子列表失败", err, nil))
		return
	}

	if err := DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("")
		})

	c.JSON(200, utils.Resp("获取成功", nil, data))
}
