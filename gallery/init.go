package main

import (
	"os"

	"github.com/McaxDev/backend/utils"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {
	var err error
	LoadConfig()
	if DB, err = utils.InitMySQL(Config.MySQL); err != nil {
		return err
	}

	if err = os.Mkdir(Config.ImagePath, 0755); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	return nil
}
