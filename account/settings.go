package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context, user *dbs.User) {

	type SettingsStruct struct {
		Name  string `json:"name"`
		Value bool   `json:"value"`
	}

	var settings []SettingsStruct
	for index, name := range utils.SettingsSlice {
		settings = append(settings, SettingsStruct{
			Name:  name,
			Value: utils.GetBitByIndex(user.Setting, uint(index)),
		})
	}

	c.JSON(200, settings)
}

func SetSetting(c *gin.Context, user *dbs.User, req struct {
	Index uint
	Value bool
}) {

	utils.UpdateBitByIndex(
		&user.Setting, req.Index, req.Value,
	)

	if err := DB.Updates(&user).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil, nil))
}
