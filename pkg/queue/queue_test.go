package queue

import (
	"email_server/models"
	"email_server/pkg/gredis"
	"email_server/pkg/logging"
	"email_server/pkg/setting"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	m.Run()
}

func Test_AMQP_PUSH(t *testing.T) {
	queue := &AMQP{
		Payload:  "11111",
		Exchange: "email_sender",
		Queue:    "email_sender",
	}
	err := queue.Push()
	if err != nil {
		t.Fatalf("测试失败: %s \n", err)
	}
}

func Test_AMQP_Pull(t *testing.T) {
	queue := &AMQP{
		Exchange: "email_sender",
		Queue:    "email_sender",
	}
	err := queue.Pull(callPull)
	if err != nil {
		t.Fatalf("测试失败: %s \n", err)
	}
}

func callPull(data string) error {
	fmt.Printf("消费数据 %s \n", data)
	return nil
}
