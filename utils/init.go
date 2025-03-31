package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(config MySQLConfig) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name,
	)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil, err
	}

	db.AutoMigrate(Tables...)

	for _, fk := range Constraints {
		if err := CreateForeignKey(db, fk); err != nil {
			return nil, err
		}
	}

	return db, nil
}
