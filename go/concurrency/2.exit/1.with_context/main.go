package main

import (
	"context"
)

type Events interface {
	DO(ctx context.Context) context.Context
}

type Signal struct {
	// Timeout  time.Duration
	// Deadline time.Time
	Cancel <-chan struct{}
}

func DoEvent(e Events, ctx context.Context) {
	e.DO(ctx)
}

func (s *Signal) DO(ctx context.Context) context.Context {
	if cancelEventFromArgCtx := s.Cancel; cancelEventFromArgCtx != nil {
		inThisFuncCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			select {
			case <-cancelEventFromArgCtx:
				cancel()
			case <-inThisFuncCtx.Done():
			}
		}()
		ctx = inThisFuncCtx
	}
	return ctx
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := Signal{}
	DoEvent(&sig, ctx) // abstract factory
	cancel()
}
