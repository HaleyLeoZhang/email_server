package util

// ----------------------------------------------------------------------
// util 测试
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
	t.Run("GetUuid", getUuid)
	t.Run("CheckEmail", checkEmail)
}

func getUuid(t *testing.T) {
	uuid := GetUuid()
	if 36 == len(uuid) {
		logging.Debug("GetUuid.success %s", uuid)
	} else {
		t.Fatalf("Create uuid fail!")
	}
}

func checkEmail(t *testing.T) {
	email := "haleyleozhang@sohu.com"
	if false == CheckEmail(email) {
		t.Fatalf("util.CheckEmail.fail")
	}
	email = "haleyleozhang-sohu.com"
	if true == CheckEmail(email) {
		t.Fatalf("util.CheckEmail.fail")
	}
}
