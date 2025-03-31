package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/yuin/goldmark"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	MD goldmark.Markdown
)

func Init() error {
	var err error
	if err := utils.LoadConfig(&Config); err != nil {
		return err
	}
	if DB, err = utils.InitMySQL(Config.MySQL); err != nil {
		return err
	}
	MD = goldmark.New()
	return nil
}
