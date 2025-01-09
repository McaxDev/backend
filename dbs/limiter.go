package dbs

import "time"

type LimiterRecord struct {
	ID     uint   `gorm:"primaryKey"`
	User   string `json:"user" gorm:"not null;type:VARCHAR(255);comment:用户"`
	Action string `json:"action" gorm:"not null;type:VARCHAR(255);comment:行为"`
	Valid  bool   `json:"-"`
	Time   time.Time
}
