package main

import (
	"fmt"
	"log"
)

func main() {

	if err := Init(); err != nil {
		log.Fatalln("初始化失败：", err.Error())
	}

	if err := GetRouter().Run(":" + Config.Port); err != nil {
		fmt.Println(err)
	}
}
