package main

import (
	"log"
	"net"

	auth "github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	Init()

	LoadConfig()

	go utils.ScheduleTask(600, func() {
		CleanMsgSent(
			&EmailSent,
			&PhoneSent,
			&QQSent,
			&QQMailSent,
		)
	})

	lis, err := net.Listen("tcp", ":"+Config.GRPCPort)
	if err != nil {
		log.Fatalln("failed to listen: " + err.Error())
	}
	defer lis.Close()

	s := grpc.NewServer()
	auth.RegisterAuthServer(s, new(AuthServer))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalln("failed to serve: " + err.Error())
		}
	}()

	r := gin.Default()

	r.GET("/captcha", SendCaptcha)
	r.GET("/email/:number", AuthEmail)
	r.GET("/phone/:number", SendPhone)
	r.GET("/qq/:number", SendQQCode)
	r.GET("/qqmail/:number", SendQQMailCode)

	if Config.SSL.Enable {
		err = r.RunTLS(
			":"+Config.HTTPPort, Config.SSL.Cert, Config.SSL.Key,
		)
	} else {
		err = r.Run(":" + Config.HTTPPort)
	}

	if err != nil {
		log.Fatalln("failed to run http server: " + err.Error())
	}
}
