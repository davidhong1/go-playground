package main01

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// 这段代码的目的就是想解耦数据结构和算法。
// 虽然使用 Strategy 算法也是可以完成的，而且比较干净，
// 但是在有些情况下，多个 Visitor 是来访问一个数据结构的不同部分，
// 这种情况下，数据结构有点像一个数据库，而各个 Visitor 会成为一个个的小应用。
// Kubectl 就是这种情况。
func Demo() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JSONVisitor)
		s.accept(XMLVisitor)
	}
}

type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

// 我们实现两个 Visitor:
// 一个是用来做 JSON 序列化的；
// 另一个是用来做 XML 序列化的。
func JSONVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XMLVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
