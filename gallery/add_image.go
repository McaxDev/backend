package main

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AddImage(c *gin.Context, user *dbs.User, _ struct{}) {

	albumID, err := strconv.Atoi(c.PostForm("album"))
	if err != nil {
		c.JSON(400, utils.Resp("请提供相册数字ID", err, nil))
		return
	}

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

	if err := DB.Model(new(dbs.Image)).First(
		"filename = ?", file.Filename,
	).Error; err == nil {
		c.JSON(400, utils.Resp("请重命名文件", err, nil))
		return
	}

	outFile, err := os.Create(filepath.Join(
		Config.ImagePath, file.Filename,
	))
	if err != nil {
		c.JSON(500, utils.Resp("无法创建文件", err, nil))
		return
	}
	defer outFile.Close()

	inFile, err := file.Open()
	if err != nil {
		c.JSON(500, utils.Resp("读取源文件失败", err, nil))
		return
	}
	defer inFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		c.JSON(500, utils.Resp("保存文件失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("文件已保存", nil, nil))
}
