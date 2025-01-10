package main

import (
	"os"

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

	if err = os.Mkdir(Config.ImagePath, 0755); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	return nil
}
