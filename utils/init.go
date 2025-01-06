package utils

import (
	auth "github.com/McaxDev/backend/auth/rpc"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	JWTKey      string
	AuthClient  auth.AuchClient
	SetMapTable = map[string]struct {
		Index   int
		Comment string
	}{
		"enableMfa":   {0, "启用MFA验证"},
		"mfaUseEmail": {1, "开启使用邮箱作为MFA方式，关闭则使用SMS"},
	}
)

const (
	LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)
