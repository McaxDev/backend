package mids

import "github.com/gin-gonic/gin"

func GetBody[T any](
	logicFunc func(c *gin.Context, req T),
) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params T
		var err error
		if c.Request.Method == "GET" {
			err = c.ShouldBindQuery(&params)
		} else {
			err = GetBodyByCtx(c, &params)
		}
		if err != nil {
			c.JSON(400, Resp("请求参数有误", err, nil))
			return
		}

		logicFunc(c, params)
	}
}
