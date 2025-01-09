package mids

import (
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
)

func BindReq[T any](
	logicFunc func(c *gin.Context, req T),
) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params T
		if err := HandleBindReq(c, &params); err != nil {
			c.AbortWithStatusJSON(400, utils.Resp("请求参数有误", err, nil))
			return
		}

		logicFunc(c, params)
	}
}

func HandleBindReq[T any](c *gin.Context, data T) error {
	if c.Request.Method == "GET" {
		return c.ShouldBindQuery(data)
	} else if MIMEIsJSON(c) {
		return utils.GetBodyByCtx(c, data)
	} else {
		return nil
	}
}
