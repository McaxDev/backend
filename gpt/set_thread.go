package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetThread(c *gin.Context, u *dbs.User, r struct {
	ID   string
	Name string
}) {

	var thread dbs.Thread
	if err := DB.First(&thread, "id = ?", r.ID).Error; err != nil {
		c.JSON(500, utils.Resp("获取会话失败", err, nil))
		return
	}

	if thread.UserID != u.ID {
		c.JSON(400, utils.Resp("此会话不属于你", nil, nil))
		return
	}

	thread.Name = r.Name
	if err := DB.Save(&thread).Error; err != nil {
		c.JSON(500, utils.Resp("会话更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("会话更新成功", nil, nil))
}
