package service

import (
	"context"
	"flag"
	"github.com/HaleyLeoZhang/email_server/conf"
	"github.com/HaleyLeoZhang/email_server/model/bo"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"os"
	"testing"
)

var (
	svr *Service
	ctx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	flag.Set("conf", "../build/app.yaml")
	if err := conf.Init(); err != nil {
		panic(err)
	}

	ctx = context.Background()
	svr = New()
	svr.Start()
	xlog.Init(conf.Conf.Log)
	os.Exit(m.Run())
}

func TestService_DoPushMessage(t *testing.T) {
	smtp := &bo.Smtp{}
	smtp.Subject = "测试"
	smtp.Body = "测试邮件发送服务"
	smtp.SenderName = "云天河官方邮件"
	smtp.ReceiverName = []string{"沐临风测试专用号"}
	smtp.Receiver = []string{"haleyleozhang@sohu.com"}
	err := svr.DoMessagePush(smtp)
	if err != nil {
		t.Fatalf("Err(%+v)", err)
	}
}
