package main

import "fmt"

type Signal struct{}

func Job() {
	fmt.Printf("Job is working...\n")
}

func ReceiveFunc(f func()) <-chan Signal {
	c := make(chan Signal)
	go func() {
		fmt.Printf("ReceiveFunc...\n")
		f()
		c <- Signal(struct{}{})
	}()
	return c
}

func main() {
	fmt.Printf("Job start..\n")
	c := ReceiveFunc(Job)
	<-c
	fmt.Printf("Job done...\n")
}
