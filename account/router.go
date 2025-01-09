package main

import (
	"github.com/McaxDev/backend/mids"
	"github.com/McaxDev/backend/utils/auth"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	authCaptcha := auth.NewAuthor(Redis).NewMid("captcha")

	r := gin.Default()
	r.Use(mids.SetJSONBodyToCtx)

	r.GET("/checkin", mids.AuthJwt(ajc, Checkin))
	r.GET("/get/blacklist", GetBlackList)
	r.GET("/get/checkin", mids.AuthJwt(ajc, GetCheckin))
	r.GET("/get/userinfo", mids.AuthJwt(ajc, GetUserInfo))
	r.GET("/get/settings", mids.AuthJwt(ajc, GetSettings))

	r.POST("/login", authCaptcha, mids.BindReq(Login))
	r.POST("/signup", authCaptcha, mids.BindReq(Signup))
	r.POST("/signout", authCaptcha, mids.AuthJwt(ajc, Signout))
	r.POST("/bindauth", mids.AuthJwt(ajc, Bind))
	r.POST("/set/settings", mids.AuthJwt(ajc, SetSettings))
	r.POST("/set/username", mids.AuthJwt(ajc, SetUsername))
	r.POST("/set/password", mids.BindReq(SetPassword))
	r.POST("/set/userinfo", mids.AuthJwt(ajc, SetUserInfo))

	ajc.OnlyAdmin = true

	r.POST("/set/blacklist", mids.AuthJwt(ajc, SetBlackList))
	r.DELETE("/del/blacklist", mids.AuthJwt(ajc, DelBlackList))

	return r
}
