package main

import (
	"context"

	"github.com/McaxDev/backend/utils/auth"
)

func (s *AuthServer) Promote(
	c context.Context, r *auth.Email,
) (*auth.Empty, error) {
	return new(auth.Empty), SendEmail(
		r.Receiver, r.Title, []byte(r.Content),
	)
}
