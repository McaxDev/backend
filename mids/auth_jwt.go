package mids

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthJwtConfig struct {
	JWTKey    string
	DB        *gorm.DB
	OnlyAdmin bool
}

func AuthJwt[T any](
	ajc AuthJwtConfig,
	logicFunc func(c *gin.Context, user *dbs.User, params T),
	preloads ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		rawToken := c.GetHeader("Authorization")
		if len(rawToken) < 8 {
			c.AbortWithStatusJSON(401, utils.Resp("token无效", nil, nil))
			return
		}

		token := rawToken[7:]

		jwtToken, err := jwt.Parse(
			token,
			func(t *jwt.Token) (any, error) {
				return []byte(ajc.JWTKey), nil
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(401, utils.Resp("token签名密钥不正确", err, nil))
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok || !jwtToken.Valid {
			c.AbortWithStatusJSON(401, utils.Resp("token格式不正确", nil, nil))
			return
		}

		userId := uint(claims["userId"].(float64))

		newToken, err := utils.GetJwt(userId, ajc.JWTKey)
		if err != nil {
			c.AbortWithStatusJSON(500, utils.Resp("生成新token失败", err, nil))
			return
		}
		c.Header("Authorization", newToken)

		user := dbs.User{Model: gorm.Model{ID: userId}}

		query := ajc.DB
		for _, value := range preloads {
			query.Preload(value)
		}

		if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(401, utils.Resp("用户不存在", nil, nil))
			return
		} else if err != nil {
			c.AbortWithStatusJSON(401, utils.Resp("查询用户失败", err, nil))
			return
		}

		if ajc.OnlyAdmin && !user.Admin {
			c.AbortWithStatusJSON(403, utils.Resp("你不是管理员", nil, nil))
			return
		}

		var params T
		if err := HandleBindReq(c, &params); err != nil {
			c.AbortWithStatusJSON(400, utils.Resp("请求参数有误", err, nil))
			return
		}

		logicFunc(c, &user, params)
	}
}
