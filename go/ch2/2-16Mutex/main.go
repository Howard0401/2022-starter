package main

import (
	"fmt"
	"sync"
	"time"
)

// Lock
type Goods struct {
	v map[string]int
	m *sync.Mutex
}

func (g *Goods) Increase(key string, in int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("Increase, Locked\n")
	g.v[key]++
	fmt.Printf("Increase Done, Unlock\n")
}

func (g *Goods) Value(key string) int {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("Value Locked\n")
	return g.v[key]
}

//Lock RWLock
type RWGoods struct {
	v map[string]int
	m *sync.RWMutex
}

func (rw *RWGoods) RVal(key string) int {
	rw.m.RLock()
	defer rw.m.RUnlock()
	return rw.v[key]
}

func (rw *RWGoods) WVal(key string, val int) {
	rw.m.Lock()
	defer rw.m.Unlock()
	rw.v[key] = val
}

func main() {
	M := &sync.Mutex{}
	g := Goods{
		v: make(map[string]int),
		m: M,
	}
	for i := 0; i < 10; i++ {
		go g.Increase("Goods1\n", i)
	}
	time.Sleep(100 * time.Microsecond)
	fmt.Println(g.Value("Goods1\n"))

	RWM := sync.RWMutex{}
	g2 := RWGoods{
		v: make(map[string]int),
		m: &RWM,
	}
	g2.WVal("key1", 1)
	fmt.Printf("g2.RVal(key1)=%v\n", g2.RVal("key1"))
}
