package main

import (
	"fmt"
	"time"
)

func GivenFood() chan string {
	ch := make(chan string)
	go func() {
		ch <- "str1"
		ch <- "str2"
		ch <- "str3"
		close(ch)
	}()
	return ch
}

func OnlyReceive(ch chan<- int) {
	fmt.Printf("OnlyReceive\n")
	go func() {
		ch <- 11
		time.Sleep(time.Second)
		fmt.Printf("OnlyReceive ch <- 11 done...\n")
		ch <- 22
		// if comment out time.Sleep(time.Second) in  OnlyRead, it maybe not be shown
		fmt.Printf("OnlyReceive ch <- 22 done...\n")
	}()
}

func OnlyRead(ch <-chan int) {
	// time.Sleep(time.Second)
	fmt.Printf("OnlyRead\n")
	// go func() {
	fmt.Printf("OnlyRead 1=%v\n", <-ch)
	fmt.Printf("OnlyRead 2=%v\n", <-ch)
	// }()
}

// FIFO First in first out
func main() {
	// ch := make(chan string)
	// ch2 := make(chan string, 6) // 可緩衝數據的容量=6
	// ch <- ""
	// <-ch2
	// close(ch)\

	ch := make(chan string)
	ch = GivenFood()

	// first way
	for {
		if name, ok := <-ch; ok {
			fmt.Println(name)
		} else {
			break
		}
	}

	// second way
	for data := range ch {
		fmt.Printf("data=%v\n", data)
	}
	// if channel contains nothing, it wouldn't send err as read.
	// write would cause error
	chInt := make(chan int)
	OnlyReceive(chInt)
	OnlyRead(chInt)
}
