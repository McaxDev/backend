package main

import (
	"encoding/json"
	"time"

	"github.com/McaxDev/Axolotland/backend/GameAPI/config"
	"github.com/McaxDev/backend/gameapi/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/docker/docker/client"
	"github.com/gorcon/rcon"
	"github.com/mholt/archiver/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Redis        *redis.Client
	Rcon         *rcon.Conn
	DockerClient *client.Client
	Compressor   *archiver.TarGz
	BindCodes    map[string]BindCodesValue
	DB           *gorm.DB
)

type BindCodesValue struct {
	Authcode string
	Expiry   time.Time
}

type RPCServer struct {
	rpc.UnimplementedGameAPIServer
}

func Init() error {
	var err error

	if DockerClient, err = client.NewClientWithOpts(
		client.WithHost(config.DockerHost),
		client.WithAPIVersionNegotiation(),
	); err != nil {
		return err
	}

	Compressor = archiver.NewTarGz()

	serversByte, err := utils.ReadFile(config.ServersPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(
		serversByte, &Servers,
	); err != nil {
		return err
	}

	return nil
}
