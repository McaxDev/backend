package main

import (
	"context"

	"github.com/McaxDev/backend/gameapi/rpc"
)

func (s *RPCServer) SendCmd(
	c context.Context, r *rpc.CmdReq,
) (*rpc.String, error) {

	srv := Servers[r.Server]

	response, err := srv.SendCmd(c, r.Cmd)
	if err != nil {
		return new(rpc.String), err
	}

	return &rpc.String{Data: response}, nil
}
