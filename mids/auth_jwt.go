package mids

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthJwt[T any](
	logicFunc LogicFunc[T], preloads ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		rawToken := c.GetHeader("Authorization")
		if len(rawToken) < 8 {
			c.JSON(401, Resp("token无效", nil, nil))
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
			c.JSON(401, Resp("token签名密钥不正确", err, nil))
			return
		}

		claims, ok := jwtToken.Claims.(*JWTClaims)
		if !ok || !jwtToken.Valid {
			c.JSON(401, Resp("token格式不正确", nil, nil))
			return
		}

		newToken, err := GetJwt(claims.UserID)
		if err != nil {
			c.JSON(500, Resp("生成新token失败", err, nil))
			return
		}
		c.Header("Authorization", newToken)

		user := dbs.User{Model: gorm.Model{ID: claims.UserID}}

		query := DB
		for _, value := range preloads {
			query.Preload(value)
		}

		if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
			c.JSON(401, Resp("用户不存在", nil, nil))
			return
		} else if err != nil {
			c.JSON(401, Resp("查询用户失败", err, nil))
			return
		}

		var params T
		if c.Request.Method == "GET" {
			err = c.ShouldBindQuery(&params)
		} else {
			err = GetBodyByCtx(c, &params)
		}
		if err != nil {
			c.JSON(400, Resp("请求参数有误", err, nil))
			return
		}

		logicFunc(&user, c, params)
	}
}
