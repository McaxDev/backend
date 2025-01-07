package dbs

import "time"

type LimiterRecord struct {
	ID     uint `gorm:"primaryKey"`
	User   string
	Action string
	Valid  bool
	Time   time.Time
}
