package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func JavaCmd(c *gin.Context, u *dbs.User, r struct {
	Server  string
	Command string
}) {

	srv := Servers[r.Server]
	if srv == nil {
		c.JSON(400, utils.Resp("不存在此服务器", nil, nil))
		return
	}

	if srv.Game != "java" {
		c.JSON(400, utils.Resp("不支持此服务器", nil, nil))
		return
	}

	conn, err := rcon.Dial(
		Config.HostAddr+":"+srv.RCON.Port, srv.RCON.Password,
	)
	if err != nil {
		c.JSON(500, utils.Resp("rcon连接建立失败", err, nil))
		return
	}
	defer conn.Close()

	resp, err := conn.Execute(r.Command)
	if err != nil {
		c.JSON(400, utils.Resp("命令执行失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("命令执行成功", nil, resp))
}
