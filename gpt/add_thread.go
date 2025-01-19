package main

import (
	"context"
	"time"

	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func AddThread(c *gin.Context, u *dbs.User, r struct {
	Model   string
	Content string
	Name    string
}) {

	assistant, exists := Models[r.Model]
	if !exists {
		c.JSON(400, utils.Resp("此模型不存在", nil, nil))
		return
	}

	ctx, canc := context.WithTimeout(
		context.Background(), 30*time.Second,
	)
	defer canc()

	if err := u.ExecWithCoins(DB, 2, false, func(tx *gorm.DB) error {

		thread, err := GPT.CreateThread(ctx, openai.ThreadRequest{
			Messages: []openai.ThreadMessage{{
				Role:    openai.ThreadMessageRoleUser,
				Content: r.Content,
			}},
		})
		if err != nil {
			return err
		}

		if err := DB.Create(&dbs.Thread{
			UserID: u.ID,
			TID:    thread.ID,
			Name:   r.Name,
		}).Error; err != nil {
			return err
		}

		run, err := GPT.CreateRun(ctx, thread.ID, openai.RunRequest{
			AssistantID: assistant,
		})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(400, utils.Resp("会话创建失败", err, nil))
		return
	}
}
