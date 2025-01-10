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

	authCaptcha := Author.NewMid("captcha")
	authEmail := Author.NewMid("email")
	authPhone := Author.NewMid("phone")

	r := gin.Default()
	r.Use(mids.SetJSONBodyToCtx)

	r.GET("/checkin", mids.AuthJwt(ajc, Checkin))
	r.GET("/get/blacklist", GetBlackList)
	r.GET("/get/checkin", mids.AuthJwt(ajc, GetCheckin))
	r.GET("/get/userinfo", mids.AuthJwt(ajc, GetUserInfo, "Album", "Comment", "Prop"))
	r.GET("/get/settings", mids.AuthJwt(ajc, GetSettings))

	r.POST("/login", mids.BindReq(Login))
	r.POST("/signup", authCaptcha, authEmail, mids.BindReq(Signup))
	r.POST("/signout", authCaptcha, authEmail, mids.AuthJwt(ajc, Signout))
	r.POST("/bind/phone", authPhone, mids.AuthJwt(ajc, BindPhone))
	r.POST("/bind/email", authEmail, mids.AuthJwt(ajc, BindEmail))
	r.POST("/set/settings", mids.AuthJwt(ajc, SetSettings))
	r.POST("/set/username", mids.AuthJwt(ajc, SetUsername))
	r.POST("/set/password", mids.BindReq(SetPassword))
	r.POST("/set/userinfo", mids.AuthJwt(ajc, SetUserInfo))

	ajc.OnlyAdmin = true

	r.POST("/set/blacklist", mids.AuthJwt(ajc, SetBlackList))
	r.DELETE("/del/blacklist", mids.AuthJwt(ajc, DelBlackList))

	return r
}
