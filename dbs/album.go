package dbs

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Cover     string  `json:"cover" gorm:"type:VARCHAR(255);comment:封面文件名"`
	Title     string  `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:标题"`
	OnlyAdmin bool    `json:"only_admin" gorm:"not null;comment:仅允许管理员"`
	GuildID   *uint   `json:"guild_id" gorm:"index;comment:公会ID"`
	Guild     Guild   `gorm:"constraint:OnDelete:SET NULL;"`
	UserID    *uint   `json:"userId" gorm:"index;comment:创建者"`
	User      User    `gorm:"constraint:OnDelete:SET NULL;"`
	Images    []Image `gorm:"constraint:OnDelete:CASCADE"`
}
