package main

import (
	account "github.com/McaxDev/backend/account/rpc"
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Edit[T any](user *account.User, c *gin.Context, req T) {

}

func HandleEdit[T any](c *gin.Context) error {

	var req dbs.Category
	if err := utils.GetBodyByCtx(c, &req); err != nil {
		return err
	}

	if err := DB.Where(
		"path = ?", req.Path,
	).Or(
		"title = ?", req.Title,
	).Or(
		"order = ?", req.Order,
	).Find(new(dbs.Category)).Error; err == nil {
		return err
	}

	return DB.Save(&req).Error
}
