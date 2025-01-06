package mids

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/gin-gonic/gin"
)

type LogicFunc[T any] func(user *dbs.User, c *gin.Context, params T)

var (
	AuthClient auth.AuthClient
)
