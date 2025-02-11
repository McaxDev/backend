package main

import (
	"time"

	"github.com/McaxDev/backend/mids"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(mids.SetJSONBodyToCtx)

	ajc := mids.AuthJwtConfig{
		JWTKey:    Config.JWTKey,
		DB:        DB,
		OnlyAdmin: false,
	}

	r.GET("/get/guild", mids.BindReq(GetGuild))
	r.GET("/get/myguild", mids.OnlyAuthJwt(ajc, GetMyGuild, "Guild", "Guild.Users"))
	r.GET("/get/guilds", GetGuilds)

	r.POST("/rename", AuthGuild(ajc, []uint{3, 4}, SetGuildName))

	r.POST("/join", AuthGuild(ajc, []uint{0}, JoinGuild))
	r.POST("/leave", OnlyAuthGuild(ajc, []uint{1, 2, 3}, Leave))

	r.POST("/review", AuthGuild(ajc, []uint{3, 4}, Review))
	r.POST("/appoint", AuthGuild(ajc, []uint{4}, Appoint))

	r.POST("/create", AuthGuild(ajc, []uint{0}, CreateGuild))
	r.POST("/transfer", AuthGuild(ajc, []uint{4}, Transfer))
	r.POST("/dissolve", OnlyAuthGuild(ajc, []uint{4}, Dissolve))

	r.POST("/donate", AuthGuild(ajc, []uint{2, 3, 4}, Donate))
	r.POST("/withdraw", OnlyAuthGuild(ajc, []uint{3, 4}, Withdraw))

	r.POST("/upgrade", OnlyAuthGuild(ajc, []uint{3, 4}, Upgrade))

	return r
}
