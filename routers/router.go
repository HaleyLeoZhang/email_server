package routers

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"email_server/routers/api/email"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api_comic := r.Group("/api/email")
	api_comic.POST("send", email.Send)
	return r
}
