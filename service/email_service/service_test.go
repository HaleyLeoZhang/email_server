package email_service

// ----------------------------------------------------------------------
// email_service 测试
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
	// t.Run("EmailDoPush", emailDoPush)
	// t.Run("CheckTmpFile", checkTmpFile)
}

func emailDoPush(t *testing.T) {
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

func checkTmpFile(t *testing.T) {
	u := &Upload{}

	result := u.CheckFile("/data/logs/app/email_server/runtime/file/file-f7739845-ab82-4319-ac2e-ade8c8a1cb8e")
	t.Fatalf("result: %v", result)
}
