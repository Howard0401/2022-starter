package main

import (
	"fmt"
	"sync"
)

func GetPipeLine(left, right int) <-chan int {
	c := make(chan int)
	go func() {
		for i := left; i < left+right; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func filterOdd(in int) (int, bool) {
	if in%2 != 0 {
		return in, false
	}
	return in, true
}

func square(in int) (int, bool) {
	return in * in, true
}

func Event(id string, num int, f func(int) (int, bool), in <-chan int) <-chan int {
	var fanOutResult []chan int
	fanOutChan := make(chan int)
	// send message to several channel
	for i := 0; i < num; i++ {
		out := make(chan int)
		go func(i int) {
			fmt.Printf("task:[%s], run goroutine num:%v\n", id, i)
			for v := range in {
				r, ok := f(v)
				if ok {
					out <- r
				}
			}
			close(out)
			fmt.Printf("task:[%s],run goroutine num:%v, done\n", id, i)
		}(i)
		fanOutResult = append(fanOutResult, out)
	}

	go func() {
		var wg sync.WaitGroup
		// fan out channal collected all fan-in channel data
		for _, v := range fanOutResult {
			wg.Add(1)
			go func(out <-chan int) {
				defer wg.Done()
				for v := range out {
					fanOutChan <- v
				}
			}(v)
		}
		wg.Wait()
		close(fanOutChan)
	}()

	return fanOutChan
}

func main() {
	in := GetPipeLine(1, 40000)
	out := Event("task1", 3, square, Event("task1", 2, filterOdd, in))
	// time.Sleep(5 * time.Second)
	for v := range out {
		fmt.Printf("v=%v\n", v)
	}
}
