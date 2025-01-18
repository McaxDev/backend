package main

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

// 备份游戏服务器的存档
func SrvBak(srv Server) error {

	srvInfo := srv.GetInfo()
	bakPath := filepath.Join(srvInfo.Path, Config.BackupFolder)

	// 检查备份目录是否存在
	if err := CreateFolder(bakPath); err != nil {
		return err
	}

	// 备份世界
	if err := Compressor.Archive(
		[]string{srvInfo.Backup.WorldPath},
		filepath.Join(
			bakPath,
			time.Now().Format("2006-01-02_15:04:05")+".tar.gz",
		),
	); err != nil {
		return err
	}

	// 获取所有备份
	files, err := os.ReadDir(bakPath)
	if err != nil {
		return err
	}

	// 将备份按时间排序
	sort.Slice(files, func(i, j int) bool {
		var timeI, timeJ time.Time
		timeI, err = time.Parse(
			"2006-01-02_15:04:05", files[i].Name()[:16],
		)
		timeJ, err = time.Parse(
			"2006-01-02_15:04:05", files[j].Name()[:16],
		)
		return timeI.After(timeJ)
	})
	if err != nil {
		return err
	}

	// 删除多余备份
	for index, file := range files {
		if index >= srvInfo.Backup.Limit {
			if err := os.Remove(
				filepath.Join(bakPath, file.Name()),
			); err != nil {
				return err
			}
		}
	}

	// 开启容器
	if err := StartContainer(srvId); err != nil {
		return err
	}

	return nil
}
