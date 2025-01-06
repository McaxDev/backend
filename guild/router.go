package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter() gin.Engine {
	r := gin.Default()

	r.Use(utils.SetBodyToCtx)

}
