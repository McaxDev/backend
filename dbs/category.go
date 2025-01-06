package dbs

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	WikiMeta
	Wiki []Wiki `json:"-"`
}
