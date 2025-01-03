package main

import (
	accountutils "github.com/McaxDev/backend/account/utils"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/get", Get)
	r.Use(utils.SetBodyToCtx)
	r.POST("/edit/:table", accountutils.AuthAdmin(Edit))
	r.POST("/delete/:table", accountutils.AuthAdmin(Delete))

	return r
}
