package main01

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func Print() {
	// 第一版，使用 interface
	// Country 和 City 都实现了同样的代码，代码都是一样的，能不能省掉呢
	c1 := Country{"China"}
	c2 := City{"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}
