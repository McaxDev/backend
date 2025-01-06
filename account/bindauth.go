package main

import (
	"context"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindAuth(user *dbs.User, c *gin.Context, req *rpc.Authcode) {

	_, err := AuthClient.Auth(
		context.Background(), req,
	)

	if err != nil {
		c.JSON(400, utils.Resp("号码验证失败", err, nil))
		return
	}

	query := DB.Model(&user)

	if req.Codetype == "telephone" {
		query = query.Update("Telephone", req.Number)
	} else {
		query = query.Update("Email", req.Number)
	}

	if err := query.Error; err != nil {
		c.JSON(500, utils.Resp("号码修改失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("号码修改成功", nil, nil))
}
