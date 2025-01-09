package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelImage(c *gin.Context, user *dbs.User, ids []uint) {

	var images []dbs.Image
	if err := DB.Where("id IN ?", ids).Find(
		&images,
	).Error; err != nil {
		c.JSON(500, utils.Resp("查找图片失败", err, nil))
		return
	}

	for _, image := range images {
		if !CheckImagePerm(user, &image) {
			c.JSON(400, utils.Resp(
				fmt.Sprint("你无权删除", image.Title),
				nil, nil,
			))
			return
		}
	}

	for _, image := range images {
		if err := DB.Transaction(func(tx *gorm.DB) error {
			if err := os.Remove(
				filepath.Join(Config.ImagePath, image.Filename),
			); err != nil {
				return err
			}
			return tx.Delete(&image).Error
		}); err != nil {
			c.JSON(500, utils.Resp(fmt.Sprintf(
				"图片%s删除失败", image.Title,
			), err, nil))
		}
	}

	c.JSON(200, utils.Resp("删除成功", nil, nil))
}
