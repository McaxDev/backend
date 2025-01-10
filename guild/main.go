package main

import (
	"log"

	"github.com/McaxDev/backend/utils"
)

func main() {

	LoadConfig()

	if err := Init(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := utils.RunGin(
		GetRouter(), Config.Port, Config.SSL,
	); err != nil {
		log.Fatalln(err.Error())
	}
}
