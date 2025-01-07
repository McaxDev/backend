package main

import (
	"sync"
	"time"

	"github.com/McaxDev/backend/auth/rpc"
	"github.com/McaxDev/backend/limiter"
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

type AuthServer struct {
	rpc.UnimplementedAuthServer
}

var (
	EmailSent  MsgSent
	PhoneSent  MsgSent
	QQSent     MsgSent
	QQMailSent MsgSent
	SMSClient  *unisms.UniSMSClient
)

func Init() {
	EmailSent = MsgSentInit()
	PhoneSent = MsgSentInit()
	QQSent = MsgSentInit()
	QQMailSent = MsgSentInit()
	SMSClient = unisms.NewClient(Config.SMS.ID, Config.SMS.Secret)
	limiter.SetRule("phone", []limiter.LimitRule{
		{Count: 1, Duration: 10 * time.Minute},
		{Count: 3, Duration: time.Hour},
		{Count: 5, Duration: 24 * time.Hour},
	})
}

func MsgSentInit() MsgSent {
	return MsgSent{
		data: make(map[string]MsgSentValue),
		lock: new(sync.RWMutex),
	}
}
