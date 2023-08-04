package main03

import (
	"fmt"
	"reflect"
	"strings"
)

// 泛型 Map-Reduce
// 简单版 Generic Map
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}

	return result
}

func Demo() {
	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}

	squaredArr := Map(nums, square)
	fmt.Println(squaredArr)

	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
}

// 因为反射是运行的事，
// 所以，如果类型出问题的话，就会有运行时的错误。
// 代码轻松编译了，但是运行时确出问题了，而且还是 panic 错误
func PanicDemo() {
	x := Map(5, 5)
	fmt.Println(x)
}
