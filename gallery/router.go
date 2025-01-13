package main

import (
	"time"

	"github.com/McaxDev/backend/mids"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/get/albums", GetAlbums)
	r.GET("/get/images", mids.BindReq(GetImages))
	r.POST("/add/album", mids.AuthJwt(ajc, AddAlbum))
	r.POST("/add/image", mids.OnlyAuthJwt(ajc, AddImage))
	r.POST("/set/album", mids.AuthJwt(ajc, SetAlbum))
	r.POST("/set/image", mids.AuthJwt(ajc, SetImage))
	r.DELETE("/del/album", mids.AuthJwt(ajc, DelAlbum))
	r.DELETE("/del/image", mids.AuthJwt(ajc, DelImage))

	return r
}
