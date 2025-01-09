package mids

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func MIMEIsJSON(c *gin.Context) bool {
	if strings.HasPrefix(
		c.GetHeader("Content-Type"),
		"application/json",
	) {
		return true
	} else {
		return false
	}
}
