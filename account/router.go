package main

import (
	"time"

	"github.com/McaxDev/backend/utils"
	u "github.com/McaxDev/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	u.InitPreloader(DB, Config.JWTKey)

	v := u.Verifier{Redis: Redis}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/checkin", u.Preload(Checkin, u.LOGIN))
	r.GET("/get/checkin", u.Preload(GetCheckin, u.LOGIN))
	r.GET("/get/myinfo", u.Preload(GetMyinfo, u.LOGIN, "Guild", "Albums", "Comments", "Props"))
	r.GET("/get/userinfo", u.Preload(GetUserinfo, u.LOGIN))
	r.GET("/get/settings", u.Preload(GetSettings, u.LOGIN))

	r.POST("/login", u.Preload(Login, u.JSON))
	r.POST("/signup", v.Auth("captcha"), v.Auth("email"), u.Preload(Signup, u.JSON))
	r.POST("/signout", v.Auth("captcha"), v.Auth("email"), u.Preload(Signout, u.LOGIN))

	r.POST("/bind/phone", v.Auth("phone"), u.Preload(BindPhone, u.LOGIN))
	r.POST("/bind/email", v.Auth("email"), u.Preload(BindEmail, u.LOGIN))
	r.POST("/unbind/phone", v.Auth("phone"), u.Preload(UnbindPhone, u.LOGIN|u.JSON))
	r.POST("/unbind/email", v.Auth("email"), u.Preload(UnbindEmail, u.LOGIN|u.JSON))

	r.POST("/set/setting", u.Preload(SetSetting, u.LOGIN))
	r.POST("/set/username", u.Preload(SetUsername, u.LOGIN))
	r.POST("/set/password", v.Auth("email"), utils.Preload(SetPassword, u.LOGIN))
	r.POST("/set/userinfo", utils.Preload(SetUserInfo, u.LOGIN))

	return r
}
