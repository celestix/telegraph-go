package tests

import (
	"testing"

	"github.com/anonyindian/telegraph-go"
)

func TestUploadPhoto01(t *testing.T) {
	path, err := telegraph.UploadFile("data/photo01.jpg")
	if err != nil {
		t.Error("Failed to upload photo01 to telegraph:", err)
		return
	} else if path == "" {
		t.Error("UploadPhoto returned empty path for photo01")
		return
	}

	t.Log("UploadFile on photo01 returned:", path)
}
