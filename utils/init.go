package utils

var (
	JWTKey      string
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
