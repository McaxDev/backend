package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func DelPosts(c *gin.Context, u *utils.User, IDs []uint) {

	var posts []utils.Post
	if err := DB.Find(&posts, "id IN ?", IDs).Error; err != nil {
		c.JSON(500, utils.Resp("查找帖子失败", err, nil))
		return
	}

	for index := range posts {
		if !CheckPostPerm(u, &posts[index]) {
			c.JSON(403, utils.Resp("权限不足", nil, nil))
			return
		}
	}

	if err := DB.Delete(&posts).Error; err != nil {
		c.JSON(500, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("删除成功", nil, nil))
}
