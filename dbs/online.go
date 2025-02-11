package dbs

import "time"

type Online struct {
	ID     uint      `json:"id" gorm:"primarykey;comment:ID"`
	Time   time.Time `json:"createdAt,omitempty" gorm:"comment:创建时间"`
	Server string    `json:"server" gorm:"type:VARCHAR(255);not null;comment:服务器"`
	Count  *int64    `json:"count" gorm:"comment:在线人数"`
}
