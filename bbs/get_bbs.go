package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBBS(c *gin.Context) {

	var forums []utils.BBS
	if err := DB.
		Preload("Forums", func(db *gorm.DB) *gorm.DB {
			return db.
				Preload("Cover", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "filename")
				}).Order("sort desc")
		}).Order("sort desc").
		Find(&forums).
		Error; err != nil {
		c.JSON(500, utils.Resp("获取论坛列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取论坛列表成功", nil, forums))
}
