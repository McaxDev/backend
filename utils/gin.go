package utils

import "github.com/gin-gonic/gin"

func RunGin(r *gin.Engine, port string, ssl SSLConfig) error {
	if ssl.Enable {
		return r.RunTLS(":"+port, ssl.Cert, ssl.Key)
	} else {
		return r.Run(":" + port)
	}
}
