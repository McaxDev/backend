package dbs

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"comment:'用户名'"`
	Password    string `json:"-" gorm:"comment:'密码'"`
	Avatar      string `gorm:"comment:'头像'"`
	Profile     string `gorm:"comment:'个人简介'"`
	Admin       bool   `gorm:"comment:'管理员'"`
	TempCoin    uint   `json:"temp_coin" gorm:"comment:'签到币'"`
	PermCoin    uint   `json:"perm_coin" gorm:"comment:'蝾螈币'"`
	Checkin     int64  `gorm:"comment:'签到记录'"`
	Setting     int64  `gorm:"comment:'设置'"`
	Email       string `gorm:"comment:'邮箱'"`
	Phone       string `gorm:"comment:'手机号'"`
	QQ          string `gorm:"comment:'QQ号'"`
	BedrockName string `gorm:"comment:'基岩版用户名'"`
	JavaName    string `gorm:"comment:'Java版用户名'"`
	Issue       []Issue
	GuildID     uint `json:"guildId" gorm:"comment:'公会ID'"`
	GuildRole   uint `json:"guildRole" gorm:"comment:'公会角色'"`
	Guild       Guild
}

type BlackList struct {
	gorm.Model
	Type   string    `gorm:"comment:'账号类型'"`
	Value  string    `gorm:"comment:'账号'"`
	Expiry time.Time `gorm:"comment:'解禁时间'"`
}
