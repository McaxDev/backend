package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetForum(c *gin.Context, _ *utils.User, r struct {
	Path  string `form:"path"`
	Page  int    `form:"page"`
	Count int    `form:"count"`
}) {

	var data utils.Forum
	if err := DB.
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("forum_id", "id", "author_id", "updated_at", "pinned", "title").
				Offset(10*(r.Page-1)).
				Limit(r.Count).
				Preload("Author", utils.LoadOwnerInfo)
		}).
		Preload("Cover", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "filename")
		}).
		Where("path = ?", r.Path).
		First(&data).
		Error; err != nil {
		c.JSON(500, utils.Resp("论坛加载失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("论坛加载成功", nil, gin.H{
		"count": DB.Model(&data).Association("Posts").Count(),
		"data":  &data,
	}))
}
