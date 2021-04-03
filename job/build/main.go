package main

import (
	"flag"
	"github.com/HaleyLeoZhang/email_server/common/constant"
	"github.com/HaleyLeoZhang/email_server/common/service"
	"github.com/HaleyLeoZhang/email_server/job/conf"
	"github.com/HaleyLeoZhang/email_server/job/http"
	"github.com/HaleyLeoZhang/go-component/driver/bootstrap"
	"github.com/HaleyLeoZhang/go-component/driver/httpserver"
	"github.com/HaleyLeoZhang/go-component/driver/xgin"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	oneService := service.New(conf.Conf, constant.SERVICE_TYPE_JOB)

	app := bootstrap.New()
	app.Start(func() { // 此部分代码，请勿阻塞进程
		oneService.Start()
		gin := xgin.New(conf.Conf.Gin)
		go httpserver.Run(conf.Conf.HttpServer, http.Init(gin))
		return
	}).Stop(func() {
		oneService.Close()
	})
}
