package utils

import (
	"os"
)

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
