package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	ip := net.ParseIP(c.Query("ip"))
	if ip == nil {
		ip = net.IP(c.ClientIP())
	}

	record, err := GeoIP.City(ip)
	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err.Error())
		return
	}

	sub := record.Subdivisions
	var province string
	if sub == nil {
		province = ""
	} else {
		province = sub[0].Names["zh-CN"]
	}

	c.JSON(200, gin.H{
		"msg":   "获取成功",
		"error": nil,
		"data": gin.H{
			"country":  record.Country.Names["zh-CN"],
			"province": province,
			"city":     record.City.Names["zh-CN"],
			"ip":       c.ClientIP(),
		},
	})
}
