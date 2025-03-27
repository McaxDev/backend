package main

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archives"
)

func Backup(c *gin.Context, u *dbs.User, id string) {

	srv := Servers[id]
	bakPath := filepath.Join(srv.Path, "axo-backup/")

	if err := Docker.ContainerStop(
		context.Background(), srv.ID, container.StopOptions{},
	); err != nil {
		c.JSON(500, utils.Resp("服务器停止失败", err, nil))
		return
	}

	if err := os.MkdirAll(bakPath, 0755); err != nil {
		c.JSON(500, utils.Resp("备份文件夹创建失败", err, nil))
		return
	}

	folders := make(map[string]string)
	for _, folder := range srv.Backup.SavePath {
		folders[filepath.Join(srv.Path, folder)] = folder
	}

	worlds, err := archives.FilesFromDisk(
		context.Background(), nil, folders,
	)
	if err != nil {
		c.JSON(500, utils.Resp("获取存档列表失败", err, nil))
		return
	}

	out, err := os.Create(
		time.Now().Format("2006-01-02_15:04:05") + ".zip",
	)
	if err != nil {
		c.JSON(500, utils.Resp("压缩文件创建失败", err, nil))
		return
	}
	defer out.Close()

	if err = (archives.CompressedArchive{
		Archival: archives.Zip{},
	}.Archive(
		context.Background(), out, worlds,
	)); err != nil {
		c.JSON(500, utils.Resp("创建压缩文件失败", err, nil))
		return
	}

	files, err := os.ReadDir(bakPath)
	if err != nil {
		c.JSON(500, utils.Resp("读取备份文件夹失败", err, nil))
		return
	}

	var errI, errJ error
	sort.Slice(files, func(i, j int) bool {
		var timeI, timeJ time.Time
		timeI, errI = time.Parse(
			"2006-01-02_15:04:05", files[i].Name()[:16],
		)
		timeJ, errJ = time.Parse(
			"2006-01-02_15:04:05", files[j].Name()[:16],
		)
		return timeI.After(timeJ)
	})
	if errI != nil || errJ != nil {
		c.JSON(500, utils.Resp("排列已有备份失败", nil, nil))
		return
	}

	for index, file := range files {
		if index >= int(srv.Backup.Limit) {
			if err := os.Remove(
				filepath.Join(bakPath, file.Name()),
			); err != nil {
				continue
			}
		}
	}

	if err := Docker.ContainerStart(
		context.Background(), srv.ID, container.StartOptions{},
	); err != nil {
		c.JSON(500, utils.Resp("服务器开启失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("备份成功", nil, nil))
}
