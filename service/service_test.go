package service

import (
	"context"
	"flag"
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xredis"
	"github.com/HaleyLeoZhang/go-component/errgroup"
	"node_puppeteer_example_go/common/conf"
	"os"
	"testing"
)

var (
	svr *Service
	ctx context.Context
)

func TestMain(m *testing.M) {
	flag.Parse()
	cfg := &conf.Config{}
	cfg.Redis = &xredis.Config{
		Name:  "local—redis",
		Proto: "tcp",
		Addr:  "192.168.56.110:6379",
		Auth:  "",
	}
	cfg.DB = &db.Config{
		Name:     "local-db",
		Type:     "mysql",
		Port:     3306,
		Database: "curl_avatar",
		User:     "",
		Password: "",
	}
	svr = New(cfg)
	ctx = context.Background()
	os.Exit(m.Run())
}


func emailDoPush(t *testing.T) {
	service := &service2.Email{}

	data := make(map[string]interface{})

	data["title"] = "测试"
	data["content"] = "文本"
	data["sender_name"] = "云天河测试"
	data["receiver"] = "229270575@qq.com,haleyleozhang@sohu.com"
	data["receiver_name"] = "沐临风,报警机器人"

	err := service.DoPush(data)
	if err != nil {
		t.Fatalf("测试失败")
	}
}

func checkTmpFile(t *testing.T) {
	u := &service.Upload{}

	result := u.CheckFile("/data/logs/app/email_server/runtime/file/file-f7739845-ab82-4319-ac2e-ade8c8a1cb8e")
	t.Fatalf("result: %v", result)
}
