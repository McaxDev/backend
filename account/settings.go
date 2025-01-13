package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context, user *dbs.User) {

	type SettingsStruct struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value bool   `json:"value"`
	}

	var settings []SettingsStruct
	for _, item := range utils.SettingsSlice {
		settings = append(settings, SettingsStruct{
			ID:    item.ID,
			Name:  item.Name,
			Value: utils.GetBitByID(user.Setting, item.ID),
		})
	}

	c.JSON(200, settings)
}

func SetSetting(c *gin.Context, user *dbs.User, req struct {
	ID    string
	Value bool
}) {

	utils.UpdateBitByID(
		&user.Setting, req.ID, req.Value,
	)

	if err := DB.Updates(&user).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil, nil))
}
