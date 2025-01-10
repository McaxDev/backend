package main

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddImage(c *gin.Context, user *dbs.User, _ struct{}) {

	albumIDInt, err := strconv.Atoi(c.PostForm("album"))
	if err != nil {
		c.JSON(400, utils.Resp("请提供相册数字ID", err, nil))
		return
	}
	albumID := uint(albumIDInt)

	var album dbs.Album
	if err := DB.Where("id = ?", albumID).Error; err != nil {
		c.JSON(500, utils.Resp("查找相册失败", err, nil))
		return
	}

	if !CheckUploadImage(user, &album) {
		c.JSON(400, utils.Resp("没有权限上传", nil, nil))
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, utils.Resp("无法读取文件", err, nil))
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	if title == "" || description == "" {
		c.JSON(400, utils.Resp("请提供标题及简介", nil, nil))
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)

	if err := DB.Transaction(func(tx *gorm.DB) error {
		outFile, err := os.Create(filepath.Join(
			Config.ImagePath, filename,
		))
		if err != nil {
			return err
		}
		defer outFile.Close()

		inFile, err := file.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}

		return tx.Create(&dbs.Image{
			Filename:    filename,
			Title:       title,
			Description: description,
			UserID:      &user.ID,
			AlbumID:     &albumID,
		}).Error
	}); err != nil {
		c.JSON(500, utils.Resp("文件上传失败", err, nil))
	}

	c.JSON(200, utils.Resp("文件已保存", nil, nil))
}
