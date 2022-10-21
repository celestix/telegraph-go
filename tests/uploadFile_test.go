package tests

import (
	"os"
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

	photo02Content, err := os.ReadFile("data/photo02.jpg")
	if err != nil {
		t.Error("Failed to load content of photo02:", err)
		return
	}

	path, err = telegraph.UploadFileByBytes(photo02Content)
	if err != nil {
		t.Error("Failed to upload photo02 to telegraph:", err)
		return
	} else if path == "" {
		t.Error("UploadPhoto returned empty path for photo02")
		return
	}
	t.Log("UploadFile on photo01 returned:", path)

}
