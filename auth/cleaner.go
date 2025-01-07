package main

import "time"

func CleanMsgSent(datas ...*MsgSent) {
	now := time.Now()
	for _, data := range datas {
		data.lock.Lock()
		for key, value := range data.data {
			if value.Expiry.Before(now) {
				delete(data.data, key)
			}
		}
		data.lock.Unlock()
	}
}
