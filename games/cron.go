package main

import "github.com/robfig/cron/v3"

func GetCron() *cron.Cron {
	c := cron.New()

	c.AddFunc("0 * * * *", SetRedisStats)

	return c
}
