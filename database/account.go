package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string  `gorm:"comment:'用户名'"`
	Password    string  `gorm:"comment:'密码'"`
	Avatar      string  `gorm:"comment:'头像'"`
	Profile     string  `gorm:"comment:'个人简介'"`
	Admin       bool    `gorm:"comment:'管理员'"`
	Money       int     `gorm:"comment:'金币'"`
	Checkin     int64   `gorm:"comment:'签到记录'"`
	Setting     int64   `gorm:"comment:'设置'"`
	Email       string  `gorm:"comment:'邮箱'"`
	Phone       string  `gorm:"comment:'手机号'"`
	QQ          string  `gorm:"comment:'QQ号'"`
	BedrockName string  `gorm:"comment:'基岩版用户名'"`
	JavaName    string  `gorm:"comment:'Java版用户名'"`
	Issue       []Issue `gorm:"foreignKey:UserID"`
}

type BlackList struct {
	gorm.Model
	Type   string    `gorm:"comment:'账号类型'"`
	Value  string    `gorm:"comment:'账号'"`
	Expiry time.Time `gorm:"comment:'解禁时间'"`
}
