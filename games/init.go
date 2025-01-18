package main

import (
	"github.com/McaxDev/backend/dbs"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {

	var err error
	if DB, err = dbs.InitDB(Config.DB); err != nil {
		return err
	}
	return nil
}
