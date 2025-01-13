package main

import (
	"log"
	"time"

	"github.com/McaxDev/backend/mids"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	LoadConfig()

	if err := Init(); err != nil {
		log.Fatalf("服务启动失败%v\n", err)
	}

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: true,
	}

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

	r.GET("/list", List)
	r.GET("/get", mids.BindReq(Get))
	r.POST("/edit", mids.AuthJwt(ajc, Edit))
	r.DELETE("/delete", mids.AuthJwt(ajc, Delete))

	if err := utils.RunGin(r, "8080", Config.SSL); err != nil {
		log.Fatalf("服务启动失败：%v\n", err)
	}
}
