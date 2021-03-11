package util

import "testing"

func TestCheckEmail(t *testing.T) {
	email := "haleyleozhang@sohu.com"
	if !CheckEmail(email) {
		t.Fatalf("failed")
	}
	email = "haleyleozhang-sohu.com"
	if CheckEmail(email) {
		t.Fatalf("failed")
	}
	t.Log("success")
}
