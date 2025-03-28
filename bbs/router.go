package main

import (
	u "github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/forums", GetForums)
	r.GET("/forum", u.Preload(GetForum, u.QUERY))
	r.GET("/post", u.Preload(GetPost, u.QUERY))

	return r
}
