package main

import (
	"os"

	"github.com/McaxDev/backend/utils"
)

var Config struct {
	AuthAddr string
	JWTKey   string
	Port     string
	DB       utils.DBConfig
	SSL      utils.SSLConfig
}

func LoadConfig() {

	Config.AuthAddr = os.Getenv("AUTH_ADDR")
	Config.JWTKey = os.Getenv("JWT_KEY")
	Config.Port = os.Getenv("PORT")

	utils.LoadDBConfig(&Config.DB)
	utils.LoadSSLConfig(&Config.SSL)
}
