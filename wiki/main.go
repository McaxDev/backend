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

	r := gin.Default()

	r.GET("/get", mids.GetBody(Get))
	r.POST("/edit", mids.AuthAdmin(Edit))
	r.DELETE("/delete", mids.AuthAdmin(Delete))

	if err := utils.RunGin(r, "8080", Config.SSL); err != nil {
		log.Fatalf("服务启动失败：%v\n", err)
	}
}
