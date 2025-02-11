package main

import (
	"context"
	"fmt"
)

type JavaSrv SrvInfo

func (js JavaSrv) GetInfo() *SrvInfo {
	si := SrvInfo(js)
	return &si
}

func (js JavaSrv) SendCmd(
	ctx context.Context, cmd string,
) (resp string, err error) {

	return SendCmdToContainer(
		fmt.Sprintf("/bin/bash rcon-cli %s", cmd),
		js.ID, ctx,
	)
}
