package main

import (
	"github.com/McaxDev/backend/mids"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	r := gin.Default()
	r.Use(mids.SetJSONBodyToCtx)
	r.POST("/add/online", mids.BindReq(AddOnline))
	r.POST("/authbind", mids.AuthJwt(ajc, AuthBind))
	r.POST("/backup", mids.AuthJwt(ajc, Backup))
	r.POST("/getbindcode", mids.BindReq(GetBindCode))
	r.POST("/getbindjava", mids.AuthJwt(ajc, GetBindJava))
	r.GET("/get/guild", mids.BindReq(GetGuild))
	r.GET("/get/mystat", mids.AuthJwt(ajc, GetMyStat))
	r.GET("/get/online", mids.BindReq(GetOnline))
	r.GET("/get/rank", mids.BindReq(GetRank))
	r.GET("/get/servers", GetServers)
	ajc.OnlyAdmin = true
	r.POST("/javacmd", mids.AuthJwt(ajc, JavaCmd))
	return r
}
