package main

import (
	"testing"
	"time"
)

func ShutdownMaker(psTime int) func(time.Duration) error {
	return func(time.Duration) error {
		time.Sleep(time.Second * time.Duration(psTime))
		return nil
	}
}

func TestConcurrencyShutdown(t *testing.T) {
	f1 := ShutdownMaker(2)
	f2 := ShutdownMaker(3)

	err := ConcurrencyShutdown(5*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err != nil {
		t.Errorf("ConcurrencyShutdown err=%s", err)
	}

	err = ConcurrencyShutdown(4*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err != nil {
		t.Errorf("ConcurrencyShutdown err=%s", err)
	}
}
