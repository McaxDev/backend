package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	Port         string `env:"PORT" def:"8080"`
	JWTKey       string `env:"JWT_KEY"`
	AccessKey    string `env:"ACCESS_KEY"`
	Docker       string `env:"DOCKER"`
	BackupFolder string
	SSL          utils.SSLConfig
	DB           dbs.DBConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
	utils.LoadSSLConfig(&Config.SSL)
	dbs.LoadDBConfig(&Config.DB)
}
