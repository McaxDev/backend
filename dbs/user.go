package dbs

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"comment:'用户名'"`
	Password    string `json:"-" gorm:"comment:'密码'"`
	Avatar      string `json:"avatar" gorm:"comment:'头像'"`
	Profile     string `json:"profile" gorm:"comment:'个人简介'"`
	Admin       bool   `json:"admin" gorm:"comment:'管理员'"`
	TempCoin    uint   `json:"tempCoin" gorm:"comment:'签到币'"`
	PermCoin    uint   `json:"permCoin" gorm:"comment:'蝾螈币'"`
	Checkin     int64  `json:"checkin" gorm:"comment:'签到记录'"`
	Setting     int64  `json:"setting" gorm:"comment:'设置'"`
	Email       string `json:"email" gorm:"comment:'邮箱'"`
	Phone       string `json:"phone" gorm:"comment:'手机号'"`
	QQ          string `json:"qq" gorm:"comment:'QQ号'"`
	BedrockName string `json:"bedrockName" gorm:"comment:'基岩版用户名'"`
	JavaName    string `json:"javaName" gorm:"comment:'Java版用户名'"`
	Issues      []Issue
	GuildID     uint `json:"guildId" gorm:"comment:'公会ID'"`
	GuildRole   uint `json:"guildRole" gorm:"comment:'公会角色'"`
	Guild       Guild
	Properties  []Property
}

func (user *User) ExecWithCoins(
	costs uint,
	permOnly bool,
	logicFunc func(tx *gorm.DB) error,
) error {
	return DB.Transaction(func(tx *gorm.DB) error {

		balance := user.PermCoin
		if !permOnly {
			balance += user.TempCoin
		}

		if balance < costs {
			return errors.New("金币不足")
		}

		if permOnly {
			user.PermCoin -= costs
		} else if user.TempCoin < costs {
			user.PermCoin -= (costs - user.TempCoin)
			user.TempCoin = 0
		} else {
			user.TempCoin -= costs
		}

		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		return logicFunc(tx)
	})
}
