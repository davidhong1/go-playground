package main01

import "fmt"

func Demo() {
	// 简单示例
	Decorator(Hello)("Hello, World!")
	// 另外一种写法，这种更清晰点
	hello := Decorator(Hello)
	hello("Hello")
}

func Decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}
