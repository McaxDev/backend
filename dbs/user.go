package dbs

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name        string     `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password    string     `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	Avatar      string     `json:"avatar" gorm:"type:VARCHAR(255);comment:头像"`
	Profile     string     `json:"profile" gorm:"type:TEXT;comment:个人简介"`
	Admin       bool       `json:"admin" gorm:"not null;comment:管理员"`
	TempCoin    uint       `json:"tempCoin" gorm:"not null;comment:签到币"`
	PermCoin    uint       `json:"permCoin" gorm:"not null;comment:蝾螈币"`
	Checkin     int64      `json:"checkin" gorm:"not null;comment:签到记录"`
	Setting     int64      `json:"setting" gorm:"not null;comment:设置"`
	Email       string     `json:"email" gorm:"type:VARCHAR(255);unique;comment:邮箱"`
	Phone       string     `json:"phone" gorm:"type:VARCHAR(255);unique;comment:手机号"`
	QQ          string     `json:"qq" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	BedrockName string     `json:"bedrockName" gorm:"type:VARCHAR(255);unique;comment:基岩版用户名"`
	JavaName    string     `json:"javaName" gorm:"type:VARCHAR(255);unique;comment:Java版用户名"`
	GuildID     *uint      `json:"guildId" gorm:"index;comment:公会ID"`
	GuildRole   uint       `json:"guildRole" gorm:"not null;comment:公会角色"`
	Guild       *Guild     `json:"guild" gorm:"constraint:OnDelete:SET NULL"`
	Props       []Property `json:"props" gorm:"constraint:OnDelete:CASCADE"`
	Comments    []Comment  `json:"comments" gorm:"constraint:OnDelete:SET NULL"`
	Albums      []Album    `json:"albums" gorm:"constraint:OnDelete:SET NULL;"`
}

func (user *User) ExecWithCoins(
	db *gorm.DB,
	costs uint,
	permOnly bool,
	logicFunc func(tx *gorm.DB) error,
) error {
	return db.Transaction(func(tx *gorm.DB) error {

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
