package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instances = []any{
	new(Guild),
	new(Wiki),
	new(Online),

	new(User),
	new(Post),
	new(Property),
	new(Album),
	new(Image),

	new(Comment),
}

func InitDB(config DBConfig) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name,
	))); err != nil {
		return nil, err
	}

	db.AutoMigrate(Instances...)

	return db, nil
}
