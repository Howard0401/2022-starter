package main_test

import (
	"sync"
	"testing"
)

var cs1 = 0
var mu1 sync.Mutex
var cs2 = 0
var mu2 sync.RWMutex

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncRWMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})
}

func BenchmarkWriteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			mu2.Lock()
			_ = cs2
			mu2.Unlock()
		}
	})
}
