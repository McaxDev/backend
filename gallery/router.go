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

	r.GET("/get/albums", GetAlbums)
	r.GET("/get/images", mids.BindReq(GetImages))
	r.POST("/add/album", mids.AuthJwt(ajc, AddAlbum))
	r.POST("/add/image", mids.AuthJwt(ajc, AddImage))
	r.POST("/edit/album", mids.AuthJwt(ajc, EditAlbum))
	r.POST("/edit/image", mids.AuthJwt(ajc, EditImage))
	r.DELETE("/del/album", mids.AuthJwt(ajc, DelAlbum))
	r.DELETE("/del/image", mids.AuthJwt(ajc, DelImage))

	return r
}
