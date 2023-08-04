package main05

import "fmt"

// 使用 go 泛型来替代 interface{} 修饰器模式
// 从例子可以看到，在某些场景下，go 泛型还是可以有效替代 interface{}，代码更易读
func Demo() {
	fmt.Println(Decorator(foo)(1, 2, 3))
	fmt.Println(Decorator(bar)("Hello", ",", "World", "!"))
}

func Decorator[T any](fn func(in ...T) T) func(...T) T {
	return func(t ...T) T {
		fmt.Println("before")
		defer fmt.Println("after")
		return fn(t...)
	}
}

func foo(s ...int) int {
	fmt.Printf("run foo %v\n", s)
	sum := 0
	for _, item := range s {
		sum += item
	}
	return sum
}

func bar(s ...string) string {
	fmt.Printf("run bar %v\n", s)
	sum := ""
	for _, item := range s {
		sum += item
	}
	return sum
}
