package main

import (
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	Port   string `env:"PORT" def:"8080"`
	JWTKey string `env:"JWT_KEY"`
	MySQL  utils.MySQLConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
}
