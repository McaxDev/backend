package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func DelComments(c *gin.Context, u *utils.User, IDs []uint) {

	var comments []utils.Comment
	if err := DB.Find(&comments, "id IN ?", IDs).Error; err != nil {
		c.JSON(500, utils.Resp("查找评论列表失败", err, nil))
		return
	}

	for index := range comments {
		if !CheckCommentPerm(u, &comments[index]) {
			c.JSON(403, utils.Resp("权限不足", nil, nil))
			return
		}
	}

	if err := DB.Delete(&comments).Error; err != nil {
		c.JSON(500, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("删除成功", nil, nil))
}
