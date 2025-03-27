package utils

import (
	"fmt"

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
