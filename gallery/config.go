package main

import (
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	JWTKey    string `env:"JWT_KEY"`
	Port      string `env:"PORT" def:"8080"`
	ImagePath string `env:"IMAGE_PATH" def:"/images"`
	MySQL     utils.MySQLConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
}
