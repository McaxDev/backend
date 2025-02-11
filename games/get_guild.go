package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
)

func GetGuild(c *gin.Context, r struct {
	Player string `form:"player"`
	Key    string `form:"key"`
}) {

	if Config.AccessKey != r.Key {
		c.Status(403)
		return
	}

	var user dbs.User
	if err := DB.Preload("Guild").First(
		&user, "bedrock_name = ?", r.Player,
	).Error; err != nil {
		c.Status(500)
		return
	}

	c.String(200, user.Guild.Name)
}
