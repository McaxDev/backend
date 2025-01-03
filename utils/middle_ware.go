package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

func SetBodyToCtx(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, Resp("无法读取请求", err, nil))
		return
	}

	c.Set("body", body)
}
