package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func Edit(c *gin.Context, user *utils.User, req utils.Wiki) {

	req.HTML = string(blackfriday.Run([]byte(req.Markdown)))

	if err := DB.Save(&req).Error; err != nil {
		c.JSON(500, utils.Resp("修改失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("修改成功", nil, nil))
}
