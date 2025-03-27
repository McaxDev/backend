package main

import (
	"log"

	"github.com/McaxDev/backend/utils"
)

func main() {

	if err := Init(); err != nil {
		log.Fatalln("初始化失败：", err.Error())
	}

	if err := utils.RunGin(
		GetRouter(), Config.Port, Config.SSL,
	); err != nil {
		log.Fatalln("启动失败：", err.Error())
	}
}
