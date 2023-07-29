package main03

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func Print() {
	// 第三版，引入 PrintStr，体现面向接口编程
	// 比如 io.Read 是个接口，主要实现 io.Read，就可以给 ioutil.ReadAll 使用
	d1 := Country{"China"}
	d2 := City{"Beijing"}
	PrintStr(d1)
	PrintStr(d2)
}
