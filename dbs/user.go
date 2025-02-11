package dbs

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name        string     `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password    string     `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	Avatar      string     `json:"avatar" gorm:"type:VARCHAR(255);comment:头像"`
	Profile     Profile    `json:"profile" gorm:"serializer:json;type:JSON;comment:个人资料"`
	Admin       bool       `json:"admin" gorm:"not null;comment:管理员"`
	TempCoin    uint       `json:"tempCoin" gorm:"not null;comment:签到币"`
	PermCoin    uint       `json:"permCoin" gorm:"not null;comment:蝾螈币"`
	Checkin     int64      `json:"-" gorm:"not null;comment:签到记录"`
	Email       string     `json:"email" gorm:"type:VARCHAR(255);unique;comment:邮箱"`
	Phone       *string    `json:"phone" gorm:"type:VARCHAR(255);unique;comment:手机号"`
	QQ          *string    `json:"qq" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	BedrockName *string    `json:"bedrockName" gorm:"type:VARCHAR(255);unique;comment:基岩版用户名"`
	JavaName    *string    `json:"javaName" gorm:"type:VARCHAR(255);unique;comment:Java版用户名"`
	GuildID     *uint      `json:"-" gorm:"index;comment:公会ID"`
	GuildRole   uint       `json:"guildRole" gorm:"not null;comment:公会角色"`
	Donation    uint       `json:"donation" gorm:"not null;comment:捐赠数额"`
	Exp         uint       `json:"exp" gorm:"not null;comment:经验值"`
	Level       uint       `json:"level" gorm:"-"`
	Setting     Setting    `json:"setting" gorm:"serializer:json;type:JSON;comment:用户设置"`
	Guild       *Guild     `json:"guild" gorm:"constraint:OnDelete:SET NULL"`
	Props       []Property `json:"props" gorm:"constraint:OnDelete:CASCADE"`
	Comments    []Comment  `json:"comments" gorm:"constraint:OnDelete:SET NULL"`
	Albums      []Album    `json:"albums" gorm:"constraint:OnDelete:SET NULL;"`
	Threads     []Thread
	TempMeta    map[string]bool `json:"-" gorm:"-"`
}

type Profile struct {
	IsMale     bool      `json:"isMale"`
	CoverImage string    `json:"coverImage"`
	Content    string    `json:"content"`
	Birthday   time.Time `json:"birthday"`
}

type Setting struct {
	Privacy  PrivacySetting
	Security SecuritySetting
}

type PrivacySetting struct {
	PubEmail    bool `json:"pubEmail" title:"公开用户资料"`
	PubPhone    bool `json:"pubPhone" title:"公开手机号"`
	PubQQ       bool `json:"pubQQ" title:"公开QQ号"`
	PubGameName bool `json:"pubGameName" title:"公开游戏名称"`
	PubGuild    bool `json:"pubGuild" title:"公开我的公会"`
	PubProps    bool `json:"pubProps" title:"公开我的道具"`
	PubComments bool `json:"pubComments" title:"公开我的评论"`
	PubAlbums   bool `json:"pubAlbums" title:"公开相册"`
	PubCoin     bool `json:"pubCoin" title:"公开金币"`
	PubGameData bool `json:"pubGameData" title:"公开游戏数据"`
	PubDonation bool `json:"pubDonation" title:"公开捐赠数额"`
	PubLevel    bool `json:"pubLevel" title:"公开我的等级"`
}

type SecuritySetting struct {
	UseMFA bool `json:"useMfa" title:"启用MFA验证"`
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

func (user *User) AfterFind(tx *gorm.DB) (err error) {

	if user.TempMeta == nil || user.TempMeta["allinfo"] == false {
		secret := "保密"
		settings := &user.Setting.Privacy

		if !settings.PubEmail {
			user.Email = secret
		}

		if !settings.PubPhone {
			user.Phone = &secret
		}

		if !settings.PubQQ {
			user.QQ = &secret
		}

		if !settings.PubGameName {
			user.JavaName = &secret
			user.BedrockName = &secret
		}

		if !settings.PubGuild {
			user.GuildID = nil
			user.GuildRole = 0
			user.Guild = nil
		}

		if !settings.PubProps {
			user.Props = nil
		}

		if !settings.PubComments {
			user.Comments = nil
		}

		if !settings.PubAlbums {
			user.Albums = nil
		}

		if !settings.PubCoin {
			user.PermCoin = 0
			user.TempCoin = 0
		}

		if !settings.PubDonation {
			user.Donation = 0
		}

		if !settings.PubLevel {
			user.Exp = 0
		} else {
			if user.Exp >= 0 {
				user.Level = 1
			} else if user.Exp >= 100 {
				user.Level = 2
			} else if user.Exp >= 500 {
				user.Level = 3
			} else if user.Exp >= 1000 {
				user.Level = 4
			} else if user.Exp >= 5000 {
				user.Level = 5
			} else if user.Exp >= 10000 {
				user.Level = 6
			}
		}
	}
	return nil
}
