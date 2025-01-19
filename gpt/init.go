package main

import (
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	GPT    *openai.Client
	Models map[string]string
)

func Init() error {
	var err error

	return nil
}
