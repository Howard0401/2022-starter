package main

import "fmt"

type Run interface {
	Running()
}

type Swim interface {
	Swimming()
}

type Sport interface {
	Run
	Swim
}

func GoSports(s Sport) {
	s.Running()
	s.Swimming()
}

type BOY struct {
	Name string
}

func (b *BOY) Running() {
	fmt.Println(b.Name + "在跑步..")
}

func (b *BOY) Swimming() {
	fmt.Println(b.Name + "在游泳...")
}

func main() {
	b := BOY{Name: "Howard"}
	GoSports(&b)
}
