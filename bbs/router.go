package main

import (
	"github.com/McaxDev/backend/mids"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(mids.SetJSONBodyToCtx)

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	r.GET("/get/post", mids.BindReq(GetPost))
	r.GET("/get/posts", mids.BindReq(GetPosts))
	r.POST("/add/post", mids.AuthJwt(ajc, AddPost))
	r.POST("/add/comment", mids.AuthJwt(ajc, AddComment))
	r.POST("/set/post", mids.AuthJwt(ajc, SetPost))
	r.POST("/set/comment", mids.AuthJwt(ajc, SetComment))
	r.DELETE("/del/posts", mids.AuthJwt(ajc, DelPosts))
	r.DELETE("/del/comments", mids.AuthJwt(ajc, DelComments))

	return r
}
