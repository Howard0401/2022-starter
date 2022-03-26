package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	DONE = "done"
)

type Signal struct {
	Int int
	Err error
}

func Job(args ...interface{}) Signal {
	if len(args) == 0 {
		return Signal{Int: -1, Err: fmt.Errorf("len(args) == 0")}
	}
	sig, ok := args[0].(Signal)
	if !ok {
		return Signal{Int: -1, Err: fmt.Errorf("sig, ok := args[0].(Signal) !ok")}
	}
	time.Sleep(time.Duration(sig.Int) * time.Second)
	return Signal{Int: sig.Int, Err: fmt.Errorf(DONE)}
}

func Event(count int, f func(args ...interface{}) Signal, args ...interface{}) chan struct{} {
	c := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer fmt.Printf("event [%v] done..\n", i)
			defer wg.Done()
			f(args...)
		}(i)
	}

	go func() {
		wg.Wait()
		// wait for all goroutines
		c <- struct{}{}
	}()

	return c
}

func main() {
	fmt.Printf("Event1 begin...\n\n")
	event1 := Event(10, Job, Signal{Int: 2})
	// <-event1 // end of signal

	timer := time.NewTimer(5 * time.Second)
	// timer := time.NewTimer(1 * time.Second) // timeout
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("timeout!!!\n")
	case <-event1:
		// fmt.Printf("event1 sig:%v...\n", sig)
		fmt.Printf("\nEvent1 done...\n")
	}
}
