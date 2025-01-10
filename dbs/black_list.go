package dbs

import (
	"time"
)

type BlackList struct {
	ID     uint      `json:"id" gorm:"primarykey;comment:ID"`
	Type   string    `json:"type" gorm:"not null;type:VARCHAR(255);comment:账号类型"`
	Value  string    `json:"value" gorm:"index;not null;type:VARCHAR(255);comment:账号"`
	Expiry time.Time `json:"expiry" gorm:"not null;comment:解禁时间"`
}
