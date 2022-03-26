package main

import (
	"fmt"
	"time"
)

func worker(arg ...interface{}) error {
	ACTION := "worker"
	if len(arg) == 0 {
		return fmt.Errorf("len(arg) == 0")
	}
	interval, ok := arg[0].(int)
	if !ok {
		fmt.Printf("%s !ok, arg=%v", ACTION, arg)
		return fmt.Errorf("interval, ok := arg[0].(int) !=ok")
	}
	time.Sleep(time.Second * (time.Duration(interval)))
	return fmt.Errorf("success")
}

func Event(f func(arg ...interface{}) error, args ...interface{}) chan error {
	c := make(chan error)
	go func() {
		c <- f(args...)
	}()
	return c
}

func main() {
	event1 := Event(worker, 2)
	fmt.Printf("Event1 begin\n")
	err := <-event1
	fmt.Printf("err=%v\n", err)
	fmt.Printf("Event1 done\n")

	event2 := Event(worker)
	fmt.Printf("Event2 begin\n")
	err = <-event2
	fmt.Printf("err=%v\n", err)

	fmt.Printf("Event2 begin\n")
}
