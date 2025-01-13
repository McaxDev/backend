package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetPassword(c *gin.Context, req struct {
	EmailID  string
	Password string
}) {

	if err := DB.Model(new(dbs.User)).Where(
		"email = ?", req.EmailID,
	).Update(
		"Password", req.Password,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, utils.Resp("不存在这个用户", nil, nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("密码修改失败", err, nil)) 
		return
	}

	c.JSON(200, utils.Resp("密码修改成功", nil, nil))
}
