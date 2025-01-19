package dbs

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name        string            `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password    string            `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	Avatar      string            `json:"avatar" gorm:"type:VARCHAR(255);comment:头像"`
	Profile     string            `json:"profile" gorm:"type:TEXT;comment:个人简介"`
	Admin       bool              `json:"admin" gorm:"not null;comment:管理员"`
	TempCoin    uint              `json:"tempCoin" gorm:"not null;comment:签到币"`
	PermCoin    uint              `json:"permCoin" gorm:"not null;comment:蝾螈币"`
	Checkin     int64             `json:"-" gorm:"not null;comment:签到记录"`
	Email       string            `json:"email" gorm:"type:VARCHAR(255);unique;comment:邮箱"`
	Phone       *string           `json:"phone" gorm:"type:VARCHAR(255);unique;comment:手机号"`
	QQ          *string           `json:"qq" gorm:"type:VARCHAR(255);unique;comment:QQ号"`
	BedrockName *string           `json:"bedrockName" gorm:"type:VARCHAR(255);unique;comment:基岩版用户名"`
	JavaName    *string           `json:"javaName" gorm:"type:VARCHAR(255);unique;comment:Java版用户名"`
	GuildID     *uint             `json:"guildId" gorm:"index;comment:公会ID"`
	GuildRole   uint              `json:"guildRole" gorm:"not null;comment:公会角色"`
	Donation    uint              `json:"donation" gorm:"not null;comment:捐赠数额"`
	Exp         uint              `json:"exp" gorm:"not null;comment:经验值"`
	Level       uint              `json:"level" gorm:"-"`
	StrMeta     map[string]string `json:"str_meta" gorm:"serializer:json;type:JSON;comment:字符串元数据"`
	BoolMeta    map[string]bool   `json:"bool_meta" gorm:"serializer:json;type:JSON;comment:布尔元数据"`
	Guild       *Guild            `json:"guild" gorm:"constraint:OnDelete:SET NULL"`
	Props       []Property        `json:"props" gorm:"constraint:OnDelete:CASCADE"`
	Comments    []Comment         `json:"comments" gorm:"constraint:OnDelete:SET NULL"`
	Albums      []Album           `json:"albums" gorm:"constraint:OnDelete:SET NULL;"`
	Threads     []Thread
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

	all, ok := tx.Statement.Context.Value("all").(bool)
	if !ok && !all {
		secret := "保密"

		if !user.BoolMeta["PubEmail"] {
			user.Email = secret
		}

		if !user.BoolMeta["PubPhone"] {
			user.Phone = &secret
		}

		if !user.BoolMeta["PubQQ"] {
			user.QQ = &secret
		}

		if !user.BoolMeta["PubGameName"] {
			user.JavaName = &secret
			user.BedrockName = &secret
		}

		if !user.BoolMeta["PubGuild"] {
			user.GuildID = nil
			user.GuildRole = 0
			user.Guild = nil
		}

		if !user.BoolMeta["PubProps"] {
			user.Props = nil
		}

		if !user.BoolMeta["PubComments"] {
			user.Comments = nil
		}

		if !user.BoolMeta["PubAlbums"] {
			user.Albums = nil
		}

		if !user.BoolMeta["PubCoin"] {
			user.PermCoin = 0
			user.TempCoin = 0
		}

		if !user.BoolMeta["PubDonation"] {
			user.Donation = 0
		}

		if !user.BoolMeta["PubExp"] {
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
