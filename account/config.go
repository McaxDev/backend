package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	JWTKey string `env:"JWT_KEY"`
	Port   string `env:"PORT" def:"8080"`
	DB     dbs.DBConfig
	Redis  utils.RedisConfig
	SSL    utils.SSLConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
	dbs.LoadDBConfig(&Config.DB)
	utils.LoadRedisConfig(&Config.Redis)
	utils.LoadSSLConfig(&Config.SSL)
}
