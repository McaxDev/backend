package main

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context, u *utils.User, r struct{}) {

	if err := DB.Select("setting").First(&u).Error; err != nil {
		c.JSON(500, utils.Resp("获取用户设置失败", err, nil))
		return
	}

	if u.Setting == nil {
		u.Setting = make(map[string]any)
	}

	c.JSON(200, utils.Resp("获取成功", nil, gin.H{
		"order":    utils.SettingOrder,
		"metadata": utils.SettingMap,
		"value":    u.Setting,
	}))
}

func SetSetting(c *gin.Context, u *utils.User, r struct {
	ID    string
	Value any
}) {

	if err := DB.Select("setting").First(&u).Error; err != nil {
		c.JSON(500, utils.Resp("获取用户设置失败", err, nil))
		return
	}

	if u.Setting == nil {
		u.Setting = make(map[string]any)
	}

	u.Setting[r.ID] = r.Value
	if err := DB.Save(&u).Error; err != nil {
		c.JSON(500, utils.Resp("设置更新失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("设置更新成功", nil, nil))
}
