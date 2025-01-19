package main

import (
	"context"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GetRank(c *gin.Context, r struct {
	Server string `form:"server"`
	Type   string `form:"type"`
	Limit  int64  `form:"limit"`
	IsAsc  bool   `form:"is_asc"`
}) {
	key := r.Server + ":" + r.Type
	var result []redis.Z
	var err error
	if r.IsAsc {
		result, err = Redis.ZRangeWithScores(
			context.Background(), key, 0, r.Limit-1,
		).Result()
	} else {
		result, err = Redis.ZRevRangeWithScores(
			context.Background(), key, 0, r.Limit-1,
		).Result()
	}
	if err != nil {
		c.JSON(500, utils.Resp("查询失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("查询成功", nil, result))
}
