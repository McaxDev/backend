package dbs

import "time"

type LimiterRecord struct {
	ID     uint      `gorm:"primaryKey"`
	User   string    `json:"user" gorm:"index;not null;type:VARCHAR(255);comment:用户"`
	Action string    `json:"action" gorm:"not null;type:VARCHAR(255);comment:行为"`
	Valid  bool      `json:"-" gorm:"-"`
	Time   time.Time `json:"time" gorm:"not null;comment:触发时间"`
}
