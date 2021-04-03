package main

import (
	"flag"
	"github.com/HaleyLeoZhang/email_server/api/conf"
	"github.com/HaleyLeoZhang/email_server/api/http"
	"github.com/HaleyLeoZhang/email_server/common/constant"
	"github.com/HaleyLeoZhang/email_server/common/service"
	"github.com/HaleyLeoZhang/go-component/driver/bootstrap"
	"github.com/HaleyLeoZhang/go-component/driver/httpserver"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	oneService := service.New(conf.Conf, constant.SERVICE_TYPE_API)

	app := bootstrap.New()
	app.Start(func() { // 此部分代码，请勿阻塞进程
		oneService.Start()
		gin := xgin.New(conf.Conf.Gin)
		go httpserver.Run(conf.Conf.HttpServer, http.Init(gin, oneService))
		return
	}).Stop(func() {
		oneService.Close()
	})
}
