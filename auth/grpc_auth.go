package main

import (
	"context"
	"errors"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/dchest/captcha"
)

func (s *RPCServer) Auth(
	c context.Context, r *rpc.Authcode,
) (*rpc.Boolean, error) {

	switch r.Codetype {
	case "email":
		success, err := AuthCode(r.Number, r.Authcode, EmailSent)
		return &rpc.Boolean{Data: success}, err
	case "phone":
		success, err := AuthCode(r.Number, r.Authcode, PhoneSent)
		return &rpc.Boolean{Data: success}, err
	case "qqmail":
		success, err := AuthCode(r.Number, r.Authcode, QQMailSent)
		return &rpc.Boolean{Data: success}, err
	case "qq":
		success, err := AuthCode(r.Number, r.Authcode, QQSent)
		return &rpc.Boolean{Data: success}, err
	case "captcha":
		if !captcha.VerifyString(r.Number, r.Authcode) {
			return &rpc.Boolean{Data: false}, nil
		}
		return &rpc.Boolean{Data: true}, nil
	default:
		return &rpc.Boolean{Data: false}, errors.New("invalid codetype")
	}
}
