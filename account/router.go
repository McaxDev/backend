package main

import (
	"time"

	"github.com/McaxDev/backend/mids"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	v := mids.Verifier{Redis: Redis}

	r := gin.Default()
	r.Use(mids.SetJSONBodyToCtx)

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/checkin", mids.OnlyAuthJwt(ajc, Checkin))
	r.GET("/get/blacklist", GetBlackList)
	r.GET("/get/checkin", mids.OnlyAuthJwt(ajc, GetCheckin))
	r.GET("/get/myinfo", mids.OnlyAuthJwt(ajc, GetMyinfo, "Guild", "Albums", "Comments", "Props"))
	r.GET("/get/userinfo", mids.BindReq(GetUserinfo))
	r.GET("/get/settings", mids.OnlyAuthJwt(ajc, GetSettings))

	r.POST("/login", mids.BindReq(Login))
	r.POST("/signup", v.Mid("captcha"), v.Mid("email"), mids.BindReq(Signup))
	r.POST("/signout", v.Mid("captcha"), v.Mid("email"), mids.OnlyAuthJwt(ajc, Signout))

	r.POST("/bind/phone", v.Mid("phone"), mids.AuthJwt(ajc, BindPhone))
	r.POST("/bind/email", v.Mid("email"), mids.AuthJwt(ajc, BindEmail))
	r.POST("/unbind/phone", v.Mid("phone"), mids.AuthJwt(ajc, UnbindPhone))
	r.POST("/unbind/email", v.Mid("email"), mids.AuthJwt(ajc, UnbindEmail))

	r.POST("/set/setting", mids.AuthJwt(ajc, SetSetting))
	r.POST("/set/username", mids.AuthJwt(ajc, SetUsername))
	r.POST("/set/password", v.Mid("email"), mids.BindReq(SetPassword))
	r.POST("/set/userinfo", mids.AuthJwt(ajc, SetUserInfo))

	ajc.OnlyAdmin = true

	r.POST("/set/blacklist", mids.AuthJwt(ajc, SetBlackList))
	r.DELETE("/del/blacklist", mids.AuthJwt(ajc, DelBlackList))

	return r
}
