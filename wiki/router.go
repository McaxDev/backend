package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/get", Get)
	r.Use(utils.SetBodyToCtx)
	r.POST("/edit/:table", utils.AuthAdmin(Edit))
	r.POST("/delete/:table", utils.AuthAdmin(Delete))

	return r
}
