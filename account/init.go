package main

import (
	"regexp"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	Redis          *redis.Client
	ChinaTime      *time.Location
	BlackListTypes []string
	isPhone        func(string) bool
	isEmail        func(string) bool
)

func Init() error {
	var err error

	if DB, err = dbs.InitDB(Config.DB); err != nil {
		return err
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host + ":" + Config.Redis.Port,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	})

	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	BlackListTypes = []string{
		"telephone", "email", "qq", "bedrock", "java",
	}

	isPhone = regexp.MustCompile(`^1\d{10}$`).MatchString
	isEmail = regexp.MustCompile(`^.+@.+\..+$`).MatchString

	return nil
}
