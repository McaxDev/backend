package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func AddPost(c *gin.Context, u *dbs.User, r struct {
	Category string
	Title    string
	Content  string
	UseMD    bool
}) {

	post := dbs.Post{
		Category: r.Category,
		Title:    r.Title,
		Source:   r.Content,
		UserID:   &u.ID,
	}

	if r.UseMD {
		post.Content = string(blackfriday.Run([]byte(r.Content)))
	} else {
		post.Content = r.Content
	}

	if err := DB.Create(&post).Error; err != nil {
		c.JSON(500, utils.Resp("发帖失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("发帖成功", nil, nil))
}
