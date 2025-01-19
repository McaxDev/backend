package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type UserCache struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	ExpiresOn string `json:"expiresOn"`
}

type Stats struct {
	Stats       map[string]map[string]int `json:"stats"`
	DataVersion int                       `json:"DataVersion"`
}

func SetRedisStats() {

	for _, server := range Servers {
		if server.Game != "java" || !server.EnableStats {
			continue
		}

		rawCaches, err := os.ReadFile(filepath.Join(
			server.Path, "usercache.json",
		))
		if err != nil {
			fmt.Println("无法读取用户缓存", err.Error())
			continue
		}

		var caches []UserCache
		if err := json.Unmarshal(rawCaches, &caches); err != nil {
			fmt.Println("用户缓存格式有误", err.Error())
			continue
		}

		for _, cache := range caches {
			rawData, err := os.ReadFile(filepath.Join(
				server.Path, "world/stats/", cache.UUID+".json",
			))
			if err != nil {
				fmt.Println("读取统计文件失败", err.Error())
				continue
			}

			var stats Stats
			if err := json.Unmarshal(rawData, &stats); err != nil {
				fmt.Println("统计文件格式有误", err.Error())
				continue
			}
			datas := stats.Stats
			if datas == nil {
				fmt.Println("无统计数据", nil)
				continue
			}

			for _, item := range []string{
				"mined", "picked_up", "crafted", "broken",
			} {
				data := datas["minecraft:"+item]
				if data == nil {
					continue
				}

				var sumed int
				for _, elem := range data {
					sumed += elem
				}

				if err := Redis.HSet(
					context.Background(), item, 1, sumed,
				).Err(); err != nil {
					continue
				}
			}

			custom := datas["minecraft:custom"]
			if custom == nil {
				continue
			}

			for _, item := range []string{
				"play_time",
				"deaths",
				"mob_kills",
				"damage_dealt",
				"drop",
			} {
				data := custom["minecraft:"+item]

				if err := Redis.HSet(
					context.Background(), item, data,
				).Err(); err != nil {
					continue
				}
			}
		}
	}
}
