package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func SetPost(c *gin.Context, u *dbs.User, r struct {
	ID       uint
	Category string
	Title    string
	Content  string
	UseMD    bool
}) {

	var post dbs.Post
	if err := DB.First(&post, "id = ?", r.ID).Error; err != nil {
		c.JSON(500, utils.Resp("查询失败", err, nil))
		return
	}

	if !CheckPostPerm(u, &post) {
		c.JSON(403, utils.Resp("你没有权限", nil, nil))
		return
	}

	if r.Category != "" {
		post.Category = r.Category
	}

	if r.Title != "" {
		post.Title = r.Title
	}

	if r.Content != post.Source {
		if r.UseMD {
			post.Content = string(blackfriday.Run([]byte(r.Content)))
		} else {
			post.Content = r.Content
		}
	}

	if err := DB.Updates(&post).Error; err != nil {
		c.JSON(500, utils.Resp("存储帖子失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("编辑帖子成功", nil, nil))
}
