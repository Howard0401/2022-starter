package main

import (
	"fmt"
	"sync"
	"time"
)

type Signal struct{}

func Job(num int, quit <-chan Signal) {
	fmt.Printf("Job %v begin...\n", num)
LOOP:
	for {
		select {
		case <-quit:
			time.Sleep(1 * time.Second)
		default:
			break LOOP
		}
	}
	fmt.Printf("Job %v done...\n", num)
}

func unbufferChanWgFunc(f func(int, <-chan Signal), num int, signal <-chan Signal) <-chan Signal {
	c := make(chan Signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("no.%v unbufferChanWgFunc...\n", i)
			f(i, signal)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Printf("start...\n")
	channel := make(chan Signal)
	msg := unbufferChanWgFunc(Job, 10, channel)

	time.Sleep(1 * time.Millisecond)
	fmt.Printf("ready to exit...\n")
	close(channel)
	<-msg
	fmt.Printf("done.\n")
}
