package main

import (
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Checkin(c *gin.Context, user *dbs.User, req struct{}) {

	iterator := time.Now().Day()

	if utils.GetBitByIndex(user.Checkin, iterator) {
		c.JSON(200, utils.Resp("你今天已经签到过啦", nil, nil))
		return
	}

	utils.UpdateBitByIndex(&user.Checkin, iterator, true)
	user.TempCoin += 1

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(500, utils.Resp("签到失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("签到成功", nil, nil))
}

func GetCheckin(c *gin.Context, user *dbs.User, req struct{}) {

	type Data struct {
		Date   int
		Status bool
	}

	var datas []Data

	for i := 1; i <= 31; i++ {
		datas = append(datas, Data{
			Date: i, Status: utils.GetBitByIndex(user.Checkin, i),
		})
	}

	c.JSON(200, utils.Resp("查询成功", nil, datas))
}
