package main

import (
	"time"

	"github.com/McaxDev/backend/dbs"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
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
