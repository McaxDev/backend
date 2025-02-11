package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewStore(client *redis.Client) RedisStore {
	return RedisStore{client: client}
}

func (s RedisStore) Set(id string, digits []byte) {

	var result string
	for _, val := range digits {
		result += strconv.Itoa(int(val))
	}

	if err := s.client.Set(
		context.Background(),
		"auth_captcha_"+id, result, 10*time.Minute,
	).Err(); err != nil {
		fmt.Printf("存储Captcha到Redis失败：%v\n", err)      
	}
}

func (s RedisStore) Get(id string, clear bool) (digits []byte) {

	val, err := s.client.Get(
		context.Background(), "auth_captcha_"+id,
	).Result()

	if err == redis.Nil {
		return nil
	} else if err != nil {
		fmt.Printf("获取CaptchaID失败：%v\n", err)  
		return nil
	}

	if clear {
		s.client.Del(context.Background(), id)
	}

	var result []byte

	for _, value := range val {
		num, _ := strconv.Atoi(string(value))
		result = append(result, byte(num))
	}

	return result
}
