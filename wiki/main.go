package main

import (
	"fmt"
	"log"
	"time"

	u "github.com/McaxDev/backend/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	LoadConfig()

	if err := Init(); err != nil {
		log.Fatalf("服务启动失败%v\n", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/wikis", GetWikis)
	r.GET("/document", u.Preload(GetDocument, u.QUERY))

	fmt.Printf("Running: %w\n", r.Run(":"+Config.Port))
}
