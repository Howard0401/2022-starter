package sync_vs_channel

import (
	"sync"
	"testing"
)

var num = 0
var mu sync.Mutex
var c = make(chan struct{}, 1)

func mockMutex() {
	mu.Lock()
	num++
	mu.Unlock()
}

func mockChanCSP() {
	c <- struct{}{}
	num++
	<-c
}

// 超級坑
// ab測試需要以Benchmark開頭的func
// ref. https://blog.wu-boy.com/2018/06/how-to-write-benchmark-in-go/
func BenchmarkTestMutex(t *testing.B) {
	for i := 0; i < t.N; i++ {
		mockMutex()
	}
}

func BenchmarkTestChanCSP(t *testing.B) {
	for i := 0; i < t.N; i++ {
		mockChanCSP()
	}
}
