package main

import (
	"fmt"
	"time"
)

func SendMsg(msg chan<- int, num int) bool {
	select {
	case msg <- num:
		return true
	default:
		return false
	}
}

func RecvMsg(msg <-chan int) (int, bool) {
	// var num int
	select {
	case num := <-msg:
		return num, true
	default:
		return 0, false
	}
}

func producer(c chan<- int) {
	i := 1
	for {
		time.Sleep(1 * time.Second)
		ok := SendMsg(c, i)
		if ok {
			fmt.Printf("[producer] send [%d] to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[producer] try to send [%d], but channel is full\n", i)
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := RecvMsg(c)
		if !ok {
			fmt.Printf("empty channel\n")
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Printf("[consumer] receive i=%v\n", i)
		if i >= 3 {
			fmt.Printf("[consumer] exit\n")
			return
		}
	}
}

func main() {
	c := make(chan int, 3)
	go producer(c)
	go consumer(c)
	select {}
}
