package utils

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AuthKind string

const (
	Email   AuthKind = "email"
	Phone   AuthKind = "phone"
	Captcha AuthKind = "captcha"
)

type Verifier struct {
	Redis *redis.Client
}

func (verifier Verifier) AuthCode(number, authcode string, kind AuthKind) error {

	account := "auth_" + string(kind) + "_" + number

	if exists, err := verifier.Redis.Exists(
		context.Background(), account,
	).Result(); err != nil {
		return err
	} else if exists == 0 {
		return errors.New(number + "没有申请验证码")
	}

	if value, err := verifier.Redis.Get(
		context.Background(), account,
	).Result(); err != nil {
		return err
	} else if value != authcode {
		return errors.New("验证码不正确")
	} else {
		return verifier.Redis.Del(context.Background(), account).Err()
	}
}

func (verifier Verifier) Auth(kind AuthKind) gin.HandlerFunc {
	return func(c *gin.Context) {

		var number string
		var code string

		switch kind {
		case Phone:
			number = c.GetHeader("X-Phone-Number")
			code = c.GetHeader("X-Phone-Code")
		case Email:
			number = c.GetHeader("X-Email-Number")
			code = c.GetHeader("X-Email-Code")
		case Captcha:
			number = c.GetHeader("X-Captcha-Id")
			code = c.GetHeader("X-Captcha-Value")
		}

		if err := verifier.AuthCode(
			number, code, kind,
		); err != nil {
			c.AbortWithStatusJSON(400,
				Resp("验证码不正确", err, nil),
			)
			return
		}
	}
}
