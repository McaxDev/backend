package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/mcstatus-io/mcutil/v4/status"
)

func SetOnline(serverId string) error {
	ctx, canc := context.WithTimeout(
		context.Background(), 10*time.Second,
	)
	defer canc()
	server := Servers[serverId]
	port, err := strconv.Atoi(server.Port)
	if err != nil {
		return err
	}
	uPort := uint16(port)
	var count *int64
	switch server.Game {
	case "bedrock":
		resp, err := status.Bedrock(ctx, Config.HostAddr, uPort)
		if err != nil {
			return err
		}
		count = resp.OnlinePlayers
	case "java":
		resp, err := status.Modern(ctx, Config.HostAddr, uPort)
		if err != nil {
			return err
		}
		count = resp.Players.Online
	default:
		return fmt.Errorf("不支持此种服务器：%w\n", server.Game)
	}
	return DB.Create(&dbs.Online{
		Time:   time.Now(),
		Server: server.ID,
		Count:  count,
	}).Error
}
