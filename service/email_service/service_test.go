package email_service

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

func Test_DoPush(t *testing.T) {
	service := &Email{}

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

// func Test_Send(t *testing.T) {
// 	service := &Smtp{
// 		Subject:      "测试",
// 		Body:         "...",
// 		Receiver:     []string{"229270575@qq.com", "haleyleozhang@sohu.com"},
// 		ReceiverName: []string{"管理员", "报警机器人"},
// 	}
// 	err := service.Send()
// 	if err != nil {
// 		t.Fatalf("邮件发送失败: %v", err)
// 	}

// }
