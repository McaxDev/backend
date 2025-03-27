package main

import (
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	JWTKey string `env:"JWT_KEY"`
	Port   string `env:"PORT" def:"8080"`
	SSL    utils.SSLConfig
	DB     utils.DBConfig
}
