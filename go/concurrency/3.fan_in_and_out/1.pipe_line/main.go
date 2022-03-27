package main

import "fmt"

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

func Event(f func(int) (int, bool), in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			r, ok := f(v)
			if ok {
				out <- r
			}
		}
		close(out)
	}()
	return out
}

func main() {
	in := GetPipeLine(1, 20)
	out := Event(square, Event(filterOdd, in))
	for v := range out {
		fmt.Printf("v=%v\n", v)
	}
}
