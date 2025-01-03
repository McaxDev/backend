package main

import (
	"encoding/json"
	"time"

	"github.com/McaxDev/backend/gameapi/rpc"
	"github.com/McaxDev/backend/utils"
	"github.com/docker/docker/client"
	"github.com/mholt/archiver/v3"
	"gorm.io/gorm"
)

var (
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
