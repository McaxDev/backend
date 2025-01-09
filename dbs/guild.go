package dbs

import "gorm.io/gorm"

type Guild struct {
	gorm.Model
	GID     string  `json:"gid" gorm:"not null;unique;type:VARCHAR(255);comment:公会ID"`
	Name    string  `json:"name" gorm:"not null;unique;type:VARCHAR(255);comment:公会名"`
	Number  uint    `json:"number" gorm:"not null;comment:公会人数"`
	Logo    string  `json:"logo" gorm:"type:VARCHAR(255);comment:LOGO路径"`
	Profile string  `json:"profile" gorm:"type:TEXT;comment:公会介绍"`
	Money   uint    `json:"money" gorm:"not null;comment:公会资金"`
	Level   uint    `json:"level" gorm:"not null;comment:公会等级"`
	Album   []Album `gorm:"constraint:OnDelete:SET NULL"`
	User    []User  `gorm:"constraint:OnDelete:SET NULL"`
}
