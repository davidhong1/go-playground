package main04

import (
	"fmt"
	"reflect"
)

// 使用 interface{} 来实现泛型
func Demo() {
	// 使用 Decorator() 时，还需要先声明一个函数签名，感觉好傻，一点都不泛型
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)
	// 如果你不想声明函数签名，就可以这样，看上去不是那么漂亮，但是 it works
	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello", "world!")
}

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})
	decoratedFunc.Set(v)

	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d\n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s\n", a, b)
	return a + b
}
