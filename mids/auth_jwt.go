package mids

import (
	"errors"

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

		user, err := HandleAuthJwt(c, ajc, preloads...)
		if err != nil {
			c.AbortWithStatusJSON(401, utils.Resp("凭证验证失败", err, nil))
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

		logicFunc(c, user, params)
	}
}

func OnlyAuthJwt(
	ajc AuthJwtConfig,
	logicFunc func(c *gin.Context, user *dbs.User),
	preloads ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := HandleAuthJwt(c, ajc, preloads...)
		if err != nil {
			c.AbortWithStatusJSON(401, utils.Resp("凭证验证失败", err, nil))
			return
		}

		if ajc.OnlyAdmin && !user.Admin {
			c.AbortWithStatusJSON(403, utils.Resp("你不是管理员", nil, nil))
			return
		}

		logicFunc(c, user)
	}
}

func HandleAuthJwt(
	c *gin.Context, ajc AuthJwtConfig, preloads ...string,
) (*dbs.User, error) {
	rawToken := c.GetHeader("Authorization")
	if len(rawToken) < 8 {
		c.AbortWithStatusJSON(401, utils.Resp("token无效", nil, nil))
		return nil, errors.New("token无效")
	}

	token := rawToken[7:]

	jwtToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (any, error) {
			return []byte(ajc.JWTKey), nil
		},
	)
	if err != nil {
		return nil, errors.New("token签名密钥不正确")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("token格式不正确")
	}

	userId := uint(claims["userId"].(float64))

	newToken, err := utils.GetJwt(userId, ajc.JWTKey)
	if err != nil {
		return nil, errors.New("生成新token失败")
	}
	c.Header("Authorization", newToken)

	user := dbs.User{}
	user.ID = userId

	query := ajc.DB.Set("all", true)
	for _, value := range preloads {
		query = query.Preload(value)
	}

	if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户不存在")
	} else if err != nil {
		return nil, errors.New("查询用户失败")
	}

	return &user, nil
}
