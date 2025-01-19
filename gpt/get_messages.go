package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context, u *dbs.User, r struct {
	TID string
}) {

	ctx, canc := context.WithTimeout(
		context.Background(), 30*time.Second,
	)
	defer canc()

	messages, err := GPT.ListMessage(
		ctx, r.TID, nil, nil, nil, nil, nil,
	)
	if err != nil {
		c.JSON(500, utils.Resp("获取消息列表失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("获取成功", nil, messages.Messages))
}
