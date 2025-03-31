package main

import (
	"github.com/McaxDev/backend/utils"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {

	var err error

	if DB, err = utils.InitMySQL(
		Config.MySQL,
	); err != nil {
		return err
	}

	return nil
}
