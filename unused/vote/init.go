package main

import (
	account "github.com/McaxDev/backend/account/rpc"
	"gorm.io/gorm"
)

var (
	AccountClient account.AccountClient
	DB            *gorm.DB
)
