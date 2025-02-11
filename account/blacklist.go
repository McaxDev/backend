package main

import (
	"time"

	"github.com/McaxDev/backend/dbs"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBlackList(c *gin.Context) {

	type NumAndExp struct {
		Value  string
		Expiry time.Time
	}
	var response map[string][]NumAndExp
	var data []dbs.BlackList
	if err := DB.Find(&data).Error; err != nil {
		c.JSON(500, utils.Resp("查找失败", err, nil))
		return
	}

	for _, value := range data {
		response[value.Type] = append(response[value.Type],
			NumAndExp{Value: value.Value, Expiry: value.Expiry},
		)
	}

	c.JSON(200, utils.Resp("查找成功", nil, response))
}

func SetBlackList(c *gin.Context, user *dbs.User, req dbs.BlackList) {

	var blackList dbs.BlackList
	if err := DB.First(
		&blackList, "type = ? AND value = ?", req.Type, req.Value,
	).Error; err == gorm.ErrRecordNotFound {
		if err := DB.Create(&req).Error; err != nil {
			c.JSON(500, utils.Resp("创建黑名单记录失败", err, nil))
			return
		}
	} else if err == nil {
		if err := DB.Model(&blackList).Where(
			"id = ?", blackList.ID,
		).Updates(&req).Error; err != nil {
			c.JSON(500, utils.Resp("更新黑名单记录失败", err, nil))
			return
		}
	} else {
		c.JSON(500, utils.Resp("查找黑名单记录失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("更新成功", nil, nil))
}

func DelBlackList(c *gin.Context, user *dbs.User, id uint) {

	if err := DB.Where("id = ?").Delete(
		new(dbs.BlackList),
	).Error; err != nil {
		c.JSON(400, utils.Resp("删除失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("删除成功", nil, nil))
}
