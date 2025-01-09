package utils

import (
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func LoadRedisConfig(config *RedisConfig) {
	config.Host = os.Getenv("REDIS_HOST")
	config.Port = GetEnv("REDIS_PORT", "6379")
	config.Password = os.Getenv("REDIS_PASSWORD")
	rawDB := os.Getenv("REDIS_DB")
	config.DB, _ = strconv.Atoi(rawDB)
}
