package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPost(c *gin.Context, _ *utils.User, r struct {
	ID     uint `form:"id"`
	Offset int  `form:"offset"`
	Limit  int  `form:"limit"`
}) {

	var post utils.Post
	if err := DB.
		Preload("Author", utils.LoadOwnerInfo).
		Preload("Reviews", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("id", "updated_at", "content", "attitude", "author_id").
				Offset(r.Offset).
				Limit(r.Limit).
				Preload("Author", utils.LoadOwnerInfo)
		}).
		First(&post, "id = ?", r.ID).
		Error; err != nil {
		c.JSON(500, utils.Resp("查询帖子失败", err, nil))
		return
	}

	var count int64
	if err := DB.
		Model(new(utils.Review)).
		Where("refer_id = ? AND refer_type = 'posts'", r.ID).
		Count(&count).
		Error; err != nil {
		c.JSON(500, utils.Resp("查询帖子评论数量失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取成功", nil, gin.H{
		"count": count,
		"data":  &post,
	}))
}
