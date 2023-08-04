package main04

import (
	"fmt"
	"reflect"
)

func Demo() {
	// 1. 用于字符串数组
	list01 := []string{"1", "2", "3", "4", "5", "6"}
	result01 := Transform(list01, func(a string) string {
		return a + a + a
	})
	fmt.Println(result01)
	// 2. 用于整形数组
	list02 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result02 := TransformInPlace(list02, func(a int) int {
		return a * 3
	})
	fmt.Println(result02)
	// 3. 用于结构体
	var list03 = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 0, 5000},
		{"Alice", 41, 0, 8000},
	}
	result03 := TransformInPlace(list03, func(e Employee) Employee {
		e.Salary += 10000
		e.Age += 1
		return e
	})
	fmt.Println(result03)

	// 4. Reduce
	var list04 = []int{1, 2, 3}
	result04 := Reduce(list04, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(result04)

	// 5. Filter
	var list05 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result05 := Filter(list05, func(a int) bool {
		return a > 3
	})
	fmt.Println(result05)
	fmt.Println(list05)
	FilterInPlace(&list05, func(a int) bool {
		return a > 3
	})
	fmt.Println(list05)
}

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

// 健壮版 Generic Map
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

// 健壮版 Generice Reduce
func Reduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

// 健壮版 Generic Filter
func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " + in.Kind().String() + " not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	// check the `slice` type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	// check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("transform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// Check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}

	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
