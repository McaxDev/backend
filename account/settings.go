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
			Value: user.BoolMeta[item.ID],
		})
	}

	c.JSON(200, settings)
}

func SetSetting(c *gin.Context, u *dbs.User, r struct {
	ID    string
	Value bool
}) {

	u.BoolMeta[r.ID] = r.Value

	if err := DB.Updates(&u).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil, nil))
}
