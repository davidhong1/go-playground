package main02

import "fmt"

type WithName struct {
	Name string
}

type Country struct {
	WithName
}

type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func Print() {
	// 引入一个叫 WithName 的结构体。这会带来另外的问题，在初始化的时候变得有点乱
	c1 := Country{WithName{"China"}}
	c2 := City{WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}
