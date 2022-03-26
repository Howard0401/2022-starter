package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(i int) {
	fmt.Printf("Job %d\n", i)
	time.Sleep(time.Second * time.Duration(i))
}

func Event(count int, f func(i int)) chan struct{} {
	quit := make(chan struct{})
	job := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			text := fmt.Sprintf("worker[%d]-", i)
			for {
				_, ok := <-job
				if !ok {
					fmt.Printf("%s done\n", text)
					return
				}
				Worker(i)
			}
		}(i)
	}
	go func() {
		<-quit
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()

	return quit
}

func main() {
	quit := Event(10, Worker)
	time.Sleep(2 * time.Second)
	// quit <- "quit"
	quit <- struct{}{}

	timer := time.NewTimer(5 * time.Second)
	// timer := time.NewTimer(1 * time.Second) // timeout
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("timeout!!!\n")
	case <-quit:
		// fmt.Printf("event1 sig:%v...\n", sig)
		fmt.Printf("\nEvent done...\n")
	}
}
