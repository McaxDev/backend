package dbs

import (
	"time"

	"gorm.io/gorm"
)

type BlackList struct {
	gorm.Model
	Type   string    `json:"type" gorm:"comment:'账号类型'"`
	Value  string    `json:"value" gorm:"comment:'账号'"`
	Expiry time.Time `json:"expiry" gorm:"comment:'解禁时间'"`
}
