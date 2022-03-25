package main

import (
	"fmt"
	"sync"
	"time"
)

type Example struct {
	Num int
	m   sync.Mutex
}

// DONT DO THIS copies lock value
func main() {
	eg := Example{Num: 0}

	go func(eg Example) {
		for {
			fmt.Println("f1...")
			eg.m.Lock()
			fmt.Println("f1 locked")
			time.Sleep(1 * time.Second)
			eg.m.Unlock()
			fmt.Println("f1 unlocked")
		}
	}(eg)
	eg.m.Lock()
	fmt.Println("f1 Lock ok")

	// copies lock value, f2 coudln't get mutex
	go func(eg Example) {
		for {
			fmt.Println("f2...")
			eg.m.Lock()
			fmt.Println("f2 locked")
			time.Sleep(3 * time.Second)
			eg.m.Unlock()
			fmt.Println("f2 unlocked")
		}
	}(eg)

	time.Sleep(20 * time.Second)
	eg.m.Unlock()
	fmt.Println("f1 unLock ok")
}
