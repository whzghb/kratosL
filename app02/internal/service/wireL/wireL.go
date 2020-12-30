package main

import (
	"fmt"
)

type MonParam string
type PlayParam string

type Mon struct {
	Name string
}

func NewMon(n MonParam) *Mon {
	return &Mon{Name: string(n)}
}

type Play struct {
	Name string
}

func NewPlay(n PlayParam) *Play {
	return &Play{Name: string(n)}
}

type Work struct {
	Player  *Play
	Monster *Mon
}

func NewWork(p *Play, m *Mon) Work {
	return Work{
		Player:  p,
		Monster: m,
	}
}

func (w Work) Start() {
	fmt.Printf("%s gogogo %s !!\n", w.Player.Name, w.Monster.Name)
}

func main()  {
	w := InitWork("aaa", "bbb")
	w.Start()
}

