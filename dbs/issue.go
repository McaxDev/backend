package dbs

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	UserID  uint   `json:"userId" gorm:"comment:用户名"`
	Title   string `json:"title" gorm:"not null;unique;type:VARCHAR(255);comment:议题标题"`
	Content string `json:"content" gorm:"type:TEXT;comment:议题内容"`
	Vote    []Vote `gorm:"constraint:OnDelete:CASCADE"`
	User    User   `gorm:"constraint:OnDelete:SET NULL"`
}
