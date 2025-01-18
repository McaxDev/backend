package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

var GeoIP *geoip2.Reader

func main() {

	var err error

	GeoIP, err = geoip2.Open("/data/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatalln("找不到数据库文件")
	}
	defer GeoIP.Close()

	r := gin.Default()
	r.GET("/get", Handler)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("服务启动失败")
	}
}
