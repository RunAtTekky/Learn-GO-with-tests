package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://blog.runat.xyz"
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"http://blog.runat.xyz",
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	want := map[string]bool{
		"https://google.com":    true,
		"https://youtube.com":   true,
		"http://blog.runat.xyz": false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	websites := make([]string, 100)

	for i := range websites {
		websites[i] = "a website"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, websites)
	}
}
