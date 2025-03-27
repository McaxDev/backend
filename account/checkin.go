package main

import (
	"time"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func Checkin(c *gin.Context, u *utils.User, r struct{}) {

	iterator := uint(time.Now().Day())

	if (u.Checkin>>iterator)&1 == 1 {
		c.JSON(200, utils.Resp("你今天已经签到过啦", nil, nil))
		return
	}

	u.Checkin |= (1 << iterator)
	u.TempCoin += 1

	if err := DB.Save(&u).Error; err != nil {
		c.JSON(500, utils.Resp("签到失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("签到成功", nil, nil))
}

func GetCheckin(c *gin.Context, u *utils.User, r struct{}) {

	type Data struct {
		Date   int  `json:"data"`
		Status bool `json:"status"`
	}

	var datas []Data

	for i := 1; i <= 31; i++ {
		datas = append(datas, Data{
			Date: i, Status: (u.Checkin>>uint(i))&1 == 1,
		})
	}

	c.JSON(200, utils.Resp("查询成功", nil, datas))
}
