package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	Port   string `env:"PORT" def:"8080"`
	JWTKey string `env:"JWT_KEY"`
	DB     dbs.DBConfig
	SSL    utils.SSLConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
	dbs.LoadDBConfig(&Config.DB)
	utils.LoadSSLConfig(&Config.SSL)
}
