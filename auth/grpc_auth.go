package main

import (
	"context"
	"errors"

	auth "github.com/McaxDev/backend/auth/rpc"
	"github.com/dchest/captcha"
)

func (s *AuthServer) Auth(
	c context.Context, r *auth.Authcode,
) (*auth.Empty, error) {

	switch r.Codetype {
	case "email":
		return new(auth.Empty), AuthCode(r.Number, r.Authcode, EmailSent)
	case "phone":
		return new(auth.Empty), AuthCode(r.Number, r.Authcode, PhoneSent)
	case "qqmail":
		return new(auth.Empty), AuthCode(r.Number, r.Authcode, QQMailSent)
	case "qq":
		return new(auth.Empty), AuthCode(r.Number, r.Authcode, QQSent)
	case "captcha":
		if !captcha.VerifyString(r.Number, r.Authcode) {
			return new(auth.Empty), errors.New("验证失败")
		}
		return new(auth.Empty), nil
	default:
		return new(auth.Empty), errors.New("invalid codetype")
	}
}
