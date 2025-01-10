package dbs

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	UserID   *uint  `json:"userId" gorm:"index;not null;comment:用户ID"`
	Property string `json:"property" gorm:"not null;comment:道具ID"`
	Count    uint   `json:"count" gorm:"not null;comment:数量"`
}
