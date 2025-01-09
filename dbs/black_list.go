package dbs

import (
	"time"

	"gorm.io/gorm"
)

type BlackList struct {
	gorm.Model
	Type   string    `json:"type" gorm:"not null;type:VARCHAR(255);comment:账号类型"`
	Value  string    `json:"value" gorm:"not null;type:VARCHAR(255);comment:账号"`
	Expiry time.Time `json:"expiry" gorm:"not null;comment:解禁时间"`
}
