package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func SetComment(c *gin.Context, u *utils.User, r struct {
	ID       uint
	Attitude *bool
	Content  string
	UseMD    bool
}) {

	comment := utils.Comment{
		Source:   r.Content,
		Attitude: r.Attitude,
		UserID:   &u.ID,
		PostID:   &r.ID,
	}

	if err := DB.First(&comment, "id = ?", r.ID).Error; err != nil {
		c.JSON(500, utils.Resp("获取评论信息失败", err, nil))
		return
	}

	if !CheckCommentPerm(u, &comment) {
		c.JSON(403, utils.Resp("权限不足", nil, nil))
		return
	}

	if r.UseMD {
		comment.Content = string(blackfriday.Run([]byte(r.Content)))
	} else {
		comment.Content = r.Content
	}

	c.JSON(200, utils.Resp("更新评论成功", nil, nil))
}
