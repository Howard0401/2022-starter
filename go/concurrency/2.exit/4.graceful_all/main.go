package main

import (
	"errors"
	"sync"
	"time"
)

type GracefulShoutdown interface {
	Shutdown(timeout time.Duration) error
}

type ShutdownerFunc func(time.Duration) error

func (f ShutdownerFunc) Shutdown(timeout time.Duration) error {
	return f(timeout)
}

func ConcurrencyShutdown(timeout time.Duration, shutdowners ...GracefulShoutdown) error { // 原來還可以這樣傳interface...
	c := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		for _, v := range shutdowners {
			wg.Add(1)
			go func(shutdowner GracefulShoutdown) {
				defer wg.Done()
				shutdowner.Shutdown(timeout)
			}(v)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("wait timeout")
	}
}

// TODO: study flow??
func SequentialShutdown(timeout time.Duration, shutdowners ...GracefulShoutdown) error { // 原來還可以這樣傳interface...
	c := make(chan struct{})
	start := time.Now()
	var left time.Duration
	timer := time.NewTimer(timeout)

	for _, v := range shutdowners {
		elasped := time.Since(start)
		left = timeout - elasped
		go func(shutdowner GracefulShoutdown) {
			shutdowner.Shutdown(left) // remains time
		}(v)
		c <- struct{}{}
	}

	timer.Reset(left)
	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("wait timeout")
	}
}

func main() {

}
