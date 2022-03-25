package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d is working..\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d done..\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				defer wg.Done()
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("woker %d start to work...\n", i)
				f(i)
			}
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func spawnGroupSignal(f func(i int), num int, gpsig *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				defer wg.Done()
				gpsig.L.Lock()
				for !ready {
					gpsig.Wait()
				}
				gpsig.L.Unlock()
				// mu.Unlock()
				fmt.Printf("woker %d start to work...\n", i)
				f(i)
			}
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func main() {
	// fmt.Println("start a group workers")
	// mu := &sync.Mutex{}
	// c := spawnGroup(worker, 5, mu)
	// time.Sleep(5 * time.Second)
	// fmt.Println("group workers start")
	// mu.Lock()
	// ready = true
	// mu.Unlock()
	// <-c

	gpsig := sync.NewCond(&sync.Mutex{})
	c2 := spawnGroupSignal(worker, 5, gpsig)
	time.Sleep(5 * time.Second)
	fmt.Println("spawnGroupSignal workers start")
	gpsig.L.Lock()
	ready = true
	gpsig.Broadcast()
	gpsig.L.Unlock()
	<-c2

}
