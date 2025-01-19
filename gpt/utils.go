package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sashabaranov/go-openai"
)

func PollRunStatus(
	gpt *openai.Client, TID, runID string,
) (msg openai.MessagesList, err error) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			run, err := gpt.RetrieveRun(
				context.Background(), TID, runID,
			)
			if err != nil {
				return nil, fmt.Errorf("无法获取执行：%w\n", err)
			}

			switch run.Status {
			case openai.RunStatusCompleted:
				return gpt.ListMessage(
					context.Background(), TID, nil, nil, nil, nil, nil,
				)
			}
		}
	}
}
