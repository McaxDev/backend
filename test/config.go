package main

import (
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	HttpPort string
	GrpcPort string
	ImageURL string
	DB       utils.DBConfig
}

func LoadConfig() {

}
