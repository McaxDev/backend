package utils

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func Resp(message string, err error, data any) gin.H {
	return gin.H{
		"message": message,
		"error":   err,
		"data":    data,
	}
}

func GetBodyByCtx(c *gin.Context, dest any) error {

	dataAny, exists := c.Get("body")
	if !exists {
		return nil
	}

	data, ok := dataAny.([]byte)
	if !ok {
		return errors.New("键body值的类型不是[]byte")
	}

	return json.Unmarshal(data, dest)
}
