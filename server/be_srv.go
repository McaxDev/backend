package main

import (
	"context"
	"fmt"
)

type BedrockSrv SrvInfo

func (bs *BedrockSrv) GetSrvInfo() *SrvInfo {
	si := SrvInfo(*bs)
	return &si
}

func (bs *BedrockSrv) SendCmd(
	ctx context.Context, cmd string,
) (string, error) {
	return SendCmdToContainer(
		fmt.Sprintf("/bin/bash send-command %s", cmd),
		bs.ID, ctx,
	)
}
