package tests

import (
	"os"
	"sync"
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
	t.Log("UploadFile on photo02 returned:", path)
}

func TestUploadPhoto02(t *testing.T) {
	for i := 0; i < 2; i++ {
		TestUploadPhoto01(t)
	}
}

func TestUploadPhotoWithWorkerPool(t *testing.T) {
	worker := func(indexes chan int, wg *sync.WaitGroup) {
		for range indexes {
			TestUploadPhoto01(t)
			wg.Done()
		}
	}
	ch := make(chan int, 2)

	var wg sync.WaitGroup
	for i := 0; i < cap(ch); i++ {
		go worker(ch, &wg)
	}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		ch <- i
	}
	wg.Wait()
	close(ch)
}
