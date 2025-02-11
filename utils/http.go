package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(message string, err error, data any) gin.H {
	var errMsg *string
	if err != nil {
		fmt.Println(message + err.Error())
		msg := err.Error()
		errMsg = &msg
	}
	return gin.H{
		"message": message,
		"error":   errMsg,
		"data":    data,
	}
}

func GetBodyByCtx(c *gin.Context, dest any) error {

	dataAny, exists := c.Get("body")
	if !exists {
		return errors.New("未知错误")
	}

	data, ok := dataAny.([]byte)
	if !ok {
		return errors.New("键body值的类型不是[]byte")
	}

	return json.Unmarshal(data, dest)
}

func Get[T any](url string) (*T, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data T
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
