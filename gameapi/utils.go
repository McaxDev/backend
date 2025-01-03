package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
)

// 开启容器
func StartContainer(name string) error {
	return DockerClient.ContainerStart(
		context.Background(), name, container.StartOptions{},
	)
}

// 关闭容器
func StopContainer(name string) error {
	return DockerClient.ContainerStop(
		context.Background(), name, container.StopOptions{},
	)
}

// 创建文件夹
func CreateFolder(path string) error {
	if result, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	} else if !result.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

func SendCmdToContainer(
	cmd, srv string, ctx context.Context,
) (string, error) {

	execute, err := DockerClient.ContainerExecCreate(
		ctx, srv, container.ExecOptions{
			Cmd:          []string{cmd},
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
	)
	if err != nil {
		return "", err
	}

	response, err := DockerClient.ContainerExecAttach(
		ctx, execute.ID, container.ExecStartOptions{Tty: true},
	)
	if err != nil {
		return "", err
	}
	defer response.Close()

	data, err := io.ReadAll(response.Reader)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
