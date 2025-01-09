package main

import (
	accountrpc "github.com/McaxDev/backend/account/rpc"
	authrpc "github.com/McaxDev/backend/auth/rpc"
	gameapirpc "github.com/McaxDev/backend/gameapi/rpc"
	"gorm.io/gorm"
)

var (
	AuthClient    authrpc.AuthClient
	AccountClient accountrpc.AccountClient
	GameapiClient gameapirpc.GameAPIClient
	DB            *gorm.DB
)
