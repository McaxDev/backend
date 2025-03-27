package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func AddComment(c *gin.Context, u *utils.User, r struct {
	PostID   uint `json:"postId" binding:"required"`
	Content  string
	UseMD    bool
	Attitude *bool
}) {

	data := utils.Comment{
		Source:   r.Content,
		Attitude: r.Attitude,
		PostID:   &r.ID,
		UserID:   &u.ID,
	}

	if r.UseMD {
		data.Content = string(blackfriday.Run([]byte(r.Content)))
	} else {
		data.Content = r.Content
	}

	if err := DB.Create(&data).Error; err != nil {
		c.JSON(500, utils.Resp("创建评论失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("发送评论成功", nil, nil))
}
