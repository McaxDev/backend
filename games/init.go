package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/docker/docker/client"
	"github.com/go-redis/redis/v8"
	"github.com/gorcon/rcon"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Redis   *redis.Client
	Rcon    *rcon.Conn
	Docker  *client.Client
	Servers = make(map[string]*Server)
)

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Game string `json:"game"`
	Path string `json:"path,omitempty"`
	RCON struct {
		Port     string
		Password string
	}
	Backup struct {
		Enable    bool
		SavePath  []string
		Frequency string
		Limit     uint
	} `json:"backup"`
	EnableStats bool
}

func Init() error {

	var err error
	if DB, err = dbs.InitDB(Config.DB); err != nil {
		return err
	}
	return nil
}
