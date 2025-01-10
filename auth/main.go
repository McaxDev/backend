package main

import (
	"log"
	"net"

	"github.com/McaxDev/backend/mids"
	"github.com/McaxDev/backend/utils"
	"github.com/McaxDev/backend/utils/auth"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	LoadConfig()

	if err := Init(); err != nil {
		log.Fatalln(err.Error())
	}

	captcha.SetCustomStore(NewStore(Redis))

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

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	r := gin.Default()

	r.GET("/captcha", SendCaptcha)
	r.GET("/email/:number", AuthEmail)
	r.GET("/phone/:number", mids.OnlyAuthJwt(ajc, SendPhone))
	r.GET("/qq/:number", SendQQCode)
	r.GET("/qqmail/:number", SendQQMailCode)

	if err := utils.RunGin(
		r, Config.HTTPPort, Config.SSL,
	); err != nil {
		log.Fatalln("failed to run http server: " + err.Error())
	}
}
