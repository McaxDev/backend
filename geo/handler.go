package main

import (
	"net"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	ip := net.ParseIP(c.Query("ip"))

	record, err := GeoIP.City(ip)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, gin.H{
		"country":  record.Country.Names["zh-CN"],
		"province": record.Subdivisions[0].Names["zh-CN"],
		"city":     record.City.Names["zh-CN"],
	})
}
