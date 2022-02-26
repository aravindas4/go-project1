package tested

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://fake.com" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"waat://fake.com",
	}

	want := map[string]bool{
		"https://google.com": true,
		"https://yahoo.com":  true,
		"https://bing.com":   true,
		"waat://fake.com":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
