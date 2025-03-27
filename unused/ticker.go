package utils

import "time"

func ScheduleTask(second uint, logicFunc func()) {
	ticker := time.NewTicker(
		time.Duration(second) * time.Second,
	)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logicFunc()
		}
	}
}
