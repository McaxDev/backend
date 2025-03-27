package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetForum(c *gin.Context, _ *utils.User, r struct {
	ID     uint `form:"id"`
	Offset int  `form:"offset"`
	Limit  int  `form:"limit"`
}) {

	var data utils.Forum
	if err := DB.
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("ID", "UpdatedAt", "Pinned", "Title").
				Offset(r.Offset).
				Limit(r.Limit).
				Preload("Author", utils.LoadUserBaseInfo)
		}).
		Where("id = ?", r.ID).
		First(&data).
		Error; err != nil {
		c.JSON(500, utils.Resp("论坛加载失败", err, nil))
		return
	}

	var count int64
	if err := DB.
		Model(new(utils.Post)).
		Where("forum_id = ?", r.ID).
		Count(&count).
		Error; err != nil {
		c.JSON(500, utils.Resp("查询帖子数量失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("论坛加载成功", nil, gin.H{
		"count": count,
		"data":  &data,
	}))
}
