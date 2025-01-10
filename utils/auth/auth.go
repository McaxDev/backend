package auth

import (
	"context"
	"errors"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Author struct {
	Redis *redis.Client
}

type CodeMsg struct {
	Number   string
	Authcode string
}

func NewAuthor(rdb *redis.Client) Author {
	return Author{Redis: rdb}
}

func (author Author) Auth(number, authcode, kind string) error {

	account := "auth_" + kind + "_" + number

	if exists, err := author.Redis.Exists(
		context.Background(), account,
	).Result(); err != nil {
		return err
	} else if exists == 0 {
		return errors.New(number + "没有申请验证码")
	}

	if value, err := author.Redis.Get(
		context.Background(), account,
	).Result(); err != nil {
		return err
	} else if value != authcode {
		return errors.New("验证码不正确")
	} else {
		return author.Redis.Del(context.Background(), account).Err()
	}
}

func (author Author) NewMid(kind string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req map[string]any
		if err := utils.GetBodyByCtx(c, &req); err != nil {
			c.AbortWithStatusJSON(400,
				utils.Resp("用户请求有误", err, nil),
			)
			return
		}

		rawNumber := req[kind+"ID"]
		rawAuthcode := req[kind+"Code"]
		number, ok1 := rawNumber.(string)
		authcode, ok2 := rawAuthcode.(string)
		if !ok1 || !ok2 {
			c.AbortWithStatusJSON(400,
				utils.Resp("验证码格式不正确", nil, nil),
			)
			return
		}

		if err := author.Auth(
			number, authcode, kind,
		); err != nil {
			c.AbortWithStatusJSON(400,
				utils.Resp("验证码不正确", err, nil),
			)
			return
		}
	}
}
