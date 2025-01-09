package main

import (
	"log"

	"github.com/McaxDev/backend/utils"
)

func main() {

	LoadConfig()

	if err := Init(); err != nil {
		log.Fatalf("初始化失败：%v\n", err)
	}

	r := GetRouter()

	if err := utils.RunGin(
		r, Config.Port, Config.SSL,
	); err != nil {
		log.Fatalf("HTTP服务器开启失败: %v\n", err)
	}
}
