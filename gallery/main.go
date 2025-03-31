package main

import (
	"fmt"
	"log"
)

func main() {

	if err := Init(); err != nil {
		log.Fatalln("初始化失败：", err.Error())
	}

	fmt.Printf("Running: %w\n", GetRouter().Run(":"+Config.Port))
}
