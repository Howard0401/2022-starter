package main

type T struct{}

func event(f func()) chan T {
	c := make(chan T)
	go func() {
		f()
	}()
	return c
}

func main() {
	c := event(func() {})
	<-c
}
