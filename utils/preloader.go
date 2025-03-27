package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type PreloaderConfig uint

var DB *gorm.DB
var JWTKey string

const (
	LOGIN PreloaderConfig = 1 << iota
	ADMIN
	BIND
	QUERY
	JSON
)

func InitPreloader(db *gorm.DB, jwtKey string) {
	DB = db
	JWTKey = jwtKey
}

func Preload[T any](
	hf func(c *gin.Context, u *User, r T), pc PreloaderConfig, preloads ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u *User
		var r T

		if pc&LOGIN != 0 {
			rawToken := c.GetHeader("Authorization")
			if len(rawToken) < 8 {
				c.AbortWithStatusJSON(400, Resp("token无效", nil, nil))
				return
			}

			token := rawToken[7:]

			jwtToken, err := jwt.Parse(
				token,
				func(t *jwt.Token) (any, error) {
					return []byte(JWTKey), nil
				},
			)
			if err != nil {
				c.AbortWithStatusJSON(400, Resp("token签名秘钥不正确", err, nil))
				return
			}

			claims, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok || !jwtToken.Valid {
				c.AbortWithStatusJSON(400, Resp("token格式不正确", nil, nil))
				return
			}

			userId := uint(claims["userId"].(float64))

			newToken, err := GetJwt(userId, JWTKey)
			if err != nil {
				c.AbortWithStatusJSON(400, Resp("生成新token失败", err, nil))
				return
			}
			c.Header("Authorization", newToken)

			user := User{ID: userId}

			query := DB
			for _, value := range preloads {
				query = query.Preload(value)
			}

			if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(400, Resp("用户不存在", err, nil))
				return
			} else if err != nil {
				c.AbortWithStatusJSON(500, Resp("查询用户失败", err, nil))
				return
			}

			u = &user
		}

		if pc&ADMIN != 0 {
			if !u.Admin {
				c.AbortWithStatusJSON(403, Resp("你不是管理员", nil, nil))
				return
			}
		}

		if pc&BIND != 0 {
			if err := c.Bind(&r); err != nil {
				c.AbortWithStatusJSON(500, Resp("请求格式有误", err, nil))
				return
			}
		}

		if pc&QUERY != 0 {
			if err := c.BindQuery(&r); err != nil {
				c.AbortWithStatusJSON(500, Resp("请求格式有误", err, nil))
				return
			}
		}

		if pc&JSON != 0 {
			if err := c.BindJSON(&r); err != nil {
				c.AbortWithStatusJSON(500, Resp("请求格式有误", err, nil))
				return
			}
		}

		hf(c, u, r)
	}

}
