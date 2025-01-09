package utils

import "os"

type AddrConfig struct {
	Host string
	Port string
}

func LoadAddrConfig(config *AddrConfig, prefix, port string) {
	config.Host = os.Getenv(prefix + "_HOST")
	if value, exists := os.LookupEnv(prefix + "_PORT"); exists {
		config.Port = value
	} else {
		config.Port = port
	}
}
