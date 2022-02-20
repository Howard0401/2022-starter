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

func main() {
	M := sync.Mutex{}
	g := Goods{
		v: make(map[string]int),
		m: &M,
	}
	for i := 0; i < 10; i++ {
		go g.Increase("Goods1\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(g.Value("Goods1\n"))

	// RWM := sync.RWMutex{}
	// g2 := RWGoods{
	// v: make(map[string]int),
	// m: &RWM,
	// }
	// ...
}
