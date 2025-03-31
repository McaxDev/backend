package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetImage(c *gin.Context, user *utils.User, req struct {
	ID          uint
	Title       string
	Description string
}) {

	var image utils.Image
	if err := DB.Preload("Album").Where("id = ?", req.ID).First(
		&image,
	).Error; err != nil {
		c.JSON(500, utils.Resp("查询图片失败", err, nil))
		return
	}

	if !CheckImagePerm(user, &image) {
		c.JSON(400, utils.Resp("你没有权限", nil, nil))
		return
	}

	if err := DB.Model(&image).Updates(&req).Error; err != nil {
		c.JSON(500, utils.Resp("图片更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("图片更新成功", nil, nil))
}
