package main

import (
	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/database"
	"github.com/McaxDev/backend/utils"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {

	if err := utils.InitDB(
		DB, Config.DB, []any{
			new(database.Wiki),
			new(database.Category),
		},
	); err != nil {
		return err
	}

	conn, err := grpc.NewClient(Config.AccountAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	utils.AccountClient = account.NewAccountClient(conn)

	return nil
}
