package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func DelThreads(c *gin.Context, u *dbs.User, IDs []uint) {

	var threads []dbs.Thread
	if err := DB.Find(&threads, "id IN ?", IDs).Error; err != nil {
		c.JSON(500, utils.Resp("查找会话失败", err, nil))
		return
	}

	for index := range threads {
		if threads[index].UserID != u.ID {
			c.JSON(400, utils.Resp("此会话不属于你", nil, nil))
			return
		}
	}

	if err := DB.Delete(&threads).Error; err != nil {
		c.JSON(500, utils.Resp("删除会话失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("删除会话成功", nil, nil))
}
