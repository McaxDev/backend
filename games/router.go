package main

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/get/player/guild", GetPlayerGuild)
	return r
}
