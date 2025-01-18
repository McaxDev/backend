package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
)

func GetPlayerGuild(c *gin.Context) {

	player := c.Query("player")
	key := c.Query("key")

	if Config.AccessKey != key {
		c.Status(403)
		return
	}

	var user dbs.User
	if err := DB.Preload("Guild").First(
		&user, "bedrock_name = ?", player,
	).Error; err != nil {
		c.Status(500)
		return
	}

	c.String(200, user.Guild.Name)
}
