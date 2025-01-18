package main

import (
	"errors"
	"regexp"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, req struct {
	Account  string
	Password string
}) {

	query := DB.Model(new(dbs.User))
	switch AccountType(req.Account) {
	case "phone":
		query = query.Where("phone = ?", req.Account)
	case "email":
		query = query.Where("email = ?", req.Account)
	default:
		query = query.Where("name = ?", req.Account)
	}

	var user dbs.User
	if err := query.First(&user).Error; errors.Is(
		err, gorm.ErrRecordNotFound,
	) {
		c.JSON(400, utils.Resp("你尚未注册", nil, nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("用户查询失败", err, nil))
		return
	}

	if req.Password != user.Password {
		c.JSON(400, utils.Resp("密码不正确", nil, nil))
		return
	}

	token, err := utils.GetJwt(user.ID, Config.JWTKey)
	if err != nil {
		c.JSON(500, utils.Resp("用户凭证生成失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("登录成功", nil, token))
}

// DetermineType determines if the input string is a phone number, email, or other
func AccountType(input string) string {
	// Regular expression for phone number (simplified for demonstration)
	phoneRegex := regexp.MustCompile(`^1[3-9][0-9]{9}$`)

	// Regular expression for email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if phoneRegex.MatchString(input) {
		return "phone"
	} else if emailRegex.MatchString(input) {
		return "email"
	}
	return "other"
}
