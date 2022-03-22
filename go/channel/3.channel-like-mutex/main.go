package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	c chan int // communication sequence process
	i int
	m sync.Mutex
	// if use mutex, we should use sync.Add(1), sync.Done(), sync.Wait(),
	// then xxx.m.Lock(), defer xxx.m.UnLock() to keep atomic, avoid race condition...
}

func Add(cnt *Counter, i int) {
	cnt.c <- i
}

func GetNewRefCounter() *Counter {
	return &Counter{
		c: make(chan int),
	}
}

func main() {
	const job = 40000
	counter := GetNewRefCounter()
	for i := 0; i < job; i++ {
		go func(counter *Counter, i int) {
			Add(counter, i)
		}(counter, i)
	}

	cnt := 0
	res := 0
	for v := range counter.c {
		cnt++
		res += v
		if job == cnt {
			break
		}
	}
	close(counter.c)
	fmt.Printf("res=%v\n", res)
}
