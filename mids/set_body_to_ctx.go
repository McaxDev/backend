package mids

import (
	"io"

	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func SetJSONBodyToCtx(c *gin.Context) {

	if !MIMEIsJSON(c) {
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, utils.Resp("无法读取请求", err, nil))
		return
	}

	c.Set("body", body)
}
