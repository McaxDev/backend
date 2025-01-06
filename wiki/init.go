package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {

	if err := utils.InitDB(
		DB, Config.DB, []any{
			new(dbs.Wiki),
			new(dbs.Category),
		},
	); err != nil {
		return err
	}

	return nil
}
