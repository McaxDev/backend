package main

import (
	"context"

	"github.com/McaxDev/backend/gameapi/rpc"
)

func (s *RPCServer) WorldBackup(
	c context.Context, r *rpc.String,
) (*rpc.Empty, error) {

	if err := SrvBak(r.Data); err != nil {
		return new(rpc.Empty), err
	}

	return new(rpc.Empty), nil
}
