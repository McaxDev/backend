package dbs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init(config DBConfig, migrates []any) error {

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

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadDBConfig(dbc *DBConfig) {

	var exists bool
	dbc.Port, exists = os.LookupEnv("DB_PORT")
	if !exists {
		dbc.Port = "3306"
	}
	dbc.Host = os.Getenv("DB_HOST")
	dbc.Name = os.Getenv("DB_NAME")
	dbc.Password = os.Getenv("DB_PASSWORD")
	dbc.User = os.Getenv("DB_USER")

}
