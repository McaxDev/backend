package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils/auth"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthServer struct {
	auth.UnimplementedAuthServer
}

var (
	SMSClient *unisms.UniSMSClient
	Redis     *redis.Client
	DB        *gorm.DB
)

func Init() error {
	var err error
	SMSClient = unisms.NewClient(Config.SMS.ID, Config.SMS.Secret)
	Redis = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host + ":" + Config.Redis.Port,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	})
	if DB, err = dbs.InitDB(Config.DB); err != nil {
		return err
	}
	return nil
}
