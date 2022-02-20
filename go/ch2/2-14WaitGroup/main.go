package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroup Add(n), n>0,  Done() Waist()
	// send by value
	s := []string{"str1", "str2", "str3", "str4", "str4"}
	var wg sync.WaitGroup
	for _, item := range s {
		wg.Add(1)
		go SayFoodName(item, &wg)
	}
	wg.Wait()
	fmt.Printf("wg.Wait() done")
}

func SayFoodName(name string, wg *sync.WaitGroup) {
	fmt.Printf("SayFoodName name=%v\n", name)
	wg.Done()
}
