package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWikis(c *gin.Context) {

	var data []utils.Wiki
	if err := DB.Preload("Documents", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "updated_at", "path", "title", "wiki_id", "sort")
	}).Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("文档站加载失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("文档站加载成功", nil, data))
}
