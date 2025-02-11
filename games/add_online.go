package main

import (
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
)

func AddOnline(c *gin.Context, r struct {
	Server    string `json:"server"`
	Count     *int64 `json:"count"`
	AccessKey string `json:"accessKey"`
}) {

	if err := c.BindJSON(&r); err != nil {
		c.Status(500)
		return
	}

	if r.AccessKey != Config.AccessKey {
		c.Status(403)
		return
	}

	if err := DB.Create(&dbs.Online{
		Time:   time.Now(),
		Server: r.Server,
		Count:  r.Count,
	}).Error; err != nil {
		c.String(500, "数据添加失败：%w", err)
		return
	}

	c.Status(200)
}
