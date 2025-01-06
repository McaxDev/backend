package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCategory(c *gin.Context) {

	var data []struct {
		dbs.WikiMeta
		Wiki []dbs.WikiMeta
	}
	if err := DB.Preload("Wiki", func(db *gorm.DB) *gorm.DB {
		return DB.Select("order", "path", "title")
	}).Model(new(dbs.Category)).Select(
		"order", "name", "title",
	).Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("目录获取失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("目录获取成功", nil, data))

}
