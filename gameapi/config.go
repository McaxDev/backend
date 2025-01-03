package main

import "os"

var Config struct {
	HTTPPort     string
	GRPCPort     string
	DockerHost   string
	ServersPath  string
	BackupFolder string
	GRPCAddr     struct {
		Account string
	}
}

func LoadConfig() {
	Config.HTTPPort = os.Getenv("HTTP_PORT")
	Config.GRPCPort = os.Getenv("GRPC_PORT")
	Config.DockerHost = os.Getenv("DOCKER_HOST")
	Config.ServersPath = os.Getenv("SERVERS_PATH")
	Config.GRPCAddr.Account = os.Getenv("GRPC_ADDR_ACCOUNT")
}
