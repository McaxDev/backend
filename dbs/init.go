package dbs

import (
	"fmt"

	"github.com/McaxDev/backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init(config utils.DBConfig, migrates []any) error {

	var err error

	if DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name,
	))); err != nil {
		return err
	}

	DB.AutoMigrate(migrates...)

	return nil
}
