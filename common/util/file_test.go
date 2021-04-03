package util

// ----------------------------------------------------------------------
// file 测试
// ----------------------------------------------------------------------
// Link  : http://www.hlzblog.top/
// GITHUB: https://github.com/HaleyLeoZhang
// ----------------------------------------------------------------------
// tesing包文档 https://golang.google.cn/pkg/testing/
// ----------------------------------------------------------------------

import (
	"testing"
)

func TestCore(t *testing.T) {
	t.Run("CheckNotExist", checkNotExist)
}

func checkNotExist(t *testing.T) {
	notExist := CheckNotExist("/data/logs/app/email_server/runtime/file/file-7e733705-1e5f-4b00-aad0-a317c50d6612")
	if false == notExist {
		t.Fatalf("result %v", notExist)
	}
}
