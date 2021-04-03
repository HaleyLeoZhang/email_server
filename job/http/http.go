package http

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) *gin.Engine {
	e.GET("/health", func(context *gin.Context) { // 服务心跳检测
		return
	})

	return e
}
