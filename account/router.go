package main

import (
	authmids "github.com/McaxDev/backend/auth/mids"
	"github.com/McaxDev/backend/mids"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/get/userinfo", mids.AuthJwt(GetUserInfo))
	router.GET("/get/settings", mids.AuthJwt(GetSettings))
	router.POST("/set/settings", mids.AuthJwt(SetSettings))
	router.POST("/set/username", mids.AuthJwt(SetUsername))
	router.POST("/set/password", mids.GetBody(SetPassword))
	router.POST("/set/userinfo", mids.AuthJwt(SetUserInfo))
	router.POST("/signup", authmids.AuthCaptcha, mids.GetBody(Signup))
	router.POST("/signout", authmids.AuthCaptcha, mids.AuthJwt(Signout))
	router.POST("/login", authmids.AuthCaptcha, mids.GetBody(Login))

	return router
}
