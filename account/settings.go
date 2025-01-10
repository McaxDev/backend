package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context, user *dbs.User) {

	type SettingsStruct struct {
		Name    string
		Comment string
		value   bool
	}

	var settings []SettingsStruct
	for name, data := range utils.SetMapTable {
		settings = append(settings, SettingsStruct{
			Name:    name,
			Comment: data.Comment,
			value:   utils.GetBitByName(user.Setting, name),
		})
	}

	c.JSON(200, settings)
}

func SetSettings(c *gin.Context, user *dbs.User, req struct {
	Name  string
	Value bool
}) {

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, utils.Resp("用户请求不正确", err, nil))
		return
	}

	utils.UpdateBitByName(
		&user.Setting, req.Name, req.Value,
	)

	if err := DB.Updates(&user).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil, nil))
}
