package http

import (
	"github.com/HaleyLeoZhang/email_server/http/email"
	"github.com/HaleyLeoZhang/email_server/service"
	"github.com/gin-gonic/gin"
)

var srv *service.Service

func Init(e *gin.Engine, srvInjection *service.Service) *gin.Engine {
	srv = srvInjection
	//e.Use() // 暂无中间件需要被设置
	{
		apiEmail := e.Group("/api/email")
		apiEmail.POST("send", email.Send)
	}

	return e
}
