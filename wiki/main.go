package main

import (
	"log"

	"github.com/McaxDev/backend/mids"
	"github.com/McaxDev/backend/utils"
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

	r.GET("/list", List)
	r.GET("/get", mids.BindReq(Get))
	r.POST("/edit", mids.AuthJwt(ajc, Edit))
	r.DELETE("/delete", mids.AuthJwt(ajc, Delete))

	if err := utils.RunGin(r, "8080", Config.SSL); err != nil {
		log.Fatalf("服务启动失败：%v\n", err)
	}
}
