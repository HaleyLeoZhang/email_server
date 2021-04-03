package http

import (
	"github.com/HaleyLeoZhang/email_server/common/service"
	"github.com/gin-gonic/gin"
)

var srv *service.Service

func Init(e *gin.Engine, srvInjection *service.Service) *gin.Engine {
	srv = srvInjection
	e.GET("/health", func(context *gin.Context) { // 服务心跳检测
		return
	})
	{
		apiEmail := e.Group("/api/email")
		apiEmail.POST("send", EmailSend)
	}

	return e
}
