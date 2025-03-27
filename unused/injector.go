package api

import "github.com/gin-gonic/gin"

func Injector(metadata string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(metadata, true)
		c.Next()
	}
}

func InjectMap(metadata gin.H) gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, value := range metadata {
			c.Set(key, value)
		}
		c.Next()
	}
}
