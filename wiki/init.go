package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/mids"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {

	var err error

	if DB, err = dbs.Init(
		&Config.DB,
	); err != nil {
		return err
	}

	mids.Init(Config.JWTKey)

	return nil
}
