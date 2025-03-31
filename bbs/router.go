package main

import (
	"time"

	u "github.com/McaxDev/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/bbs", GetBBS)
	r.GET("/forum", u.Preload(GetForum, u.QUERY))
	r.GET("/post", u.Preload(GetPost, u.QUERY))

	return r
}
