package queue

// ----------------------------------------------------------------------
// queue 测试
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// tesing包文档 https://golang.google.cn/pkg/testing/
// ----------------------------------------------------------------------

import (
	"email_server/models"
	"email_server/pkg/gredis"
	"email_server/pkg/logging"
	"email_server/pkg/setting"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	m.Run()
}

func TestCore(t *testing.T) {
	t.Run("Push", Push)
	t.Run("Pull", Pull)
}

func Push(t *testing.T) {
	q := GetEmailQueue()
	q.SetPayload([]byte(`{"title":"For test","content":"unit test","sender_name":"local","receiver":"haleyleozhang@sohu.com","receiver_name":""}`))

	err := q.Push()
	if err != nil {
		t.Fatalf("生产消息失败: %s \n", err)
	}
}

var testFlagClose chan int

func Pull(t *testing.T) {
	testFlagClose = make(chan int, 1)
	q := GetEmailQueue()
	go func() {
		err := q.Pull(callPull)
		if err != nil {
			t.Fatalf("消费消息失败: %s \n", err)
		}
	}()
	<-testFlagClose
	q.Close() // 记得单元测试关闭连接
}

func callPull(payload []byte) error {
	logging.Debug("消费数据 %v \n", string(payload))
	testFlagClose <- 1
	return nil
}
