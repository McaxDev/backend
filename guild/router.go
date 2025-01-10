package main

import (
	"github.com/McaxDev/backend/mids"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	r.GET("/get/guild", mids.BindReq(GetGuild))
	r.GET("/get/myguild", mids.OnlyAuthJwt(ajc, GetMyGuild))
	r.GET("/get/guilds", GetGuilds)

	r.Use(mids.SetJSONBodyToCtx)

	return r
}
