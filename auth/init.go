package main

import (
	"sync"
	"time"

	"github.com/McaxDev/backend/auth/rpc"
	misc "github.com/McaxDev/backend/misc/rpc"
	unisms "github.com/apistd/uni-go-sdk/sms"
)

type MsgSent struct {
	data map[string]MsgSentValue
	lock *sync.RWMutex
}

type MsgSentValue struct {
	Expiry   time.Time
	Authcode string
}

type RPCServer struct {
	rpc.UnimplementedAuthServer
}

var (
	EmailSent  MsgSent
	PhoneSent  MsgSent
	QQSent     MsgSent
	QQMailSent MsgSent
	SMSClient  *unisms.UniSMSClient
	MiscClient misc.MiscClient
)

func Init() {
	EmailSent = MsgSentInit()
	PhoneSent = MsgSentInit()
	QQSent = MsgSentInit()
	QQMailSent = MsgSentInit()
	SMSClient = unisms.NewClient(config.SMS.ID, config.SMS.Secret)
}

func MsgSentInit() MsgSent {
	return MsgSent{
		data: make(map[string]MsgSentValue),
		lock: new(sync.RWMutex),
	}
}
