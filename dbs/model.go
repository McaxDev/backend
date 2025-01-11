package dbs

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
