package main

import (
	"context"

	"github.com/McaxDev/backend/auth/rpc"
)

func (s *AuthServer) Promote(
	c context.Context, r *rpc.Email,
) (*rpc.Empty, error) {
	return new(rpc.Empty), SendEmail(
		r.Receiver, r.Title, []byte(r.Content),
	)
}
