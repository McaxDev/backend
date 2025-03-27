package main

import (
	"bytes"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context, u *utils.User, r struct {
	ForumID uint   `json:"forumId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UseMD   bool   `json:"useMd"`
}) {

	var html string
	if r.UseMD {
		var buffer bytes.Buffer
		if err := MD.Convert([]byte(r.Content), &buffer); err != nil {
			c.JSON(500, utils.Resp("渲染Markdown失败", err, nil))
			return
		}
		html = buffer.String()
	} else {
		html = r.Content
	}

	if err := DB.Create(&utils.Post{
		ForumID:  &r.ForumID,
		Title:    r.Title,
		Markdown: r.Content,
		HTML:     html,
	}).Error; err != nil {
		c.JSON(500, utils.Resp("发帖失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("发帖成功", nil, nil))
}
