package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func GetCron() *cron.Cron {
	c := cron.New()

	c.AddFunc("0 * * * *", func() {
		SetRedisStats()
		if err := SetOnline("paper"); err != nil {
			fmt.Printf("存储paper在线信息失败：%w\n", err)
		}
		if err := SetOnline("bedrock"); err != nil {
			fmt.Printf("存储bedrock在线信息失败%w\n", err)
		}
	})

	return c
}
