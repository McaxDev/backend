package main

import (
	"os"

	"github.com/McaxDev/backend/utils"
)

var Config struct {
	GRPCPort   string
	HTTPPort   string
	GeoSrvAddr string
	SSL        utils.SSLConfig
	SMS        struct {
		ID        string
		Secret    string
		Signature string
		Template  string
	}
	SMTP struct {
		Server   string
		Port     string
		Mail     string
		Password string
	}
}

func LoadConfig() {
	Config.GRPCPort = os.Getenv("GRPC_PORT")
	Config.HTTPPort = os.Getenv("HTTP_PORT")
	Config.SMS.ID = os.Getenv("SMS_ID")
	Config.SMS.Secret = os.Getenv("SMS_SECRET")
	Config.SMS.Signature = os.Getenv("SMS_SIGNATURE")
	Config.SMS.Template = os.Getenv("SMS_TEMPLATE")
	Config.SMTP.Server = os.Getenv("SMTP_SERVER")
	Config.SMTP.Port = os.Getenv("SMTP_PORT")
	Config.SMTP.Mail = os.Getenv("SMTP_MAIL")
	Config.SMTP.Password = os.Getenv("SMTP_PASSWORD")
	utils.LoadSSLConfig(&Config.SSL)
}
