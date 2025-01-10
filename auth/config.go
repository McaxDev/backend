package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
)

var Config struct {
	JWTKey     string `env:"JWT_KEY"`
	GRPCPort   string `env:"GRPC_PORT" def:"50051"`
	HTTPPort   string `env:"HTTP_PORT" def:"8080"`
	GeoSrvAddr string `env:"GEO_SRV_ADDR"`
	SMS        struct {
		ID        string `env:"SMS_ID"`
		Secret    string `env:"SMS_SECRET"`
		Signature string `env:"SMS_SIGNATURE"`
		Template  string `env:"SMS_TEMPLATE"`
	}
	SMTP struct {
		Server   string `env:"SMTP_SERVER"`
		Port     string `env:"SMTP_PORT"`
		Mail     string `env:"SMTP_MAIL"`
		Password string `env:"SMTP_PASSWORD"`
	}
	SSL   utils.SSLConfig
	DB    dbs.DBConfig
	Redis utils.RedisConfig
}

func LoadConfig() {
	utils.LoadConfig(&Config)
	utils.LoadConfig(&Config.SMS)
	utils.LoadConfig(&Config.SMTP)
	utils.LoadSSLConfig(&Config.SSL)
	dbs.LoadDBConfig(&Config.DB)
	utils.LoadRedisConfig(&Config.Redis)
}
