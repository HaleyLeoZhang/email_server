package util

import "testing"

func TestGetUuid(t *testing.T) {
	uuid := GetUuid()
	if 36 == len(uuid) {
		t.Logf("GetUuid.success %v", uuid)
	} else {
		t.Fatalf("Create uuid fail!")
	}
}
