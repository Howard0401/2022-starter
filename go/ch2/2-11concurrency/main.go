package main

import (
	"fmt"
	"runtime"
	"time"
)

func ShowBook() {
	fmt.Println("ShowBook")
}

// go run -race main.go
func main() {
	// go ShowBook()

	for i := 0; i < 10; i++ {
		// go func() {
		// fmt.Println(fmt.Sprintf("i=%d", i))
		// }()
		go func(j int) {
			fmt.Println(fmt.Sprintf("i=%d", j))
		}(i)
	}
	runtime.Gosched()

	time.Sleep(time.Second * 1)
}
