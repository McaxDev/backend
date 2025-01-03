package main

import (
	"regexp"
	"time"

	"github.com/McaxDev/backend/account/rpc"
	auth "github.com/McaxDev/backend/auth/rpc"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	ChinaTime      *time.Location
	BlackListTypes []string
	AuthClient     auth.AuthClient
	isPhone        func(string) bool
	isEmail        func(string) bool
)

type RPCServer struct {
	rpc.UnimplementedAccountServer
}

func Init() error {
	var err error

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
