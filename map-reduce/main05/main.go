package main05

import "fmt"

// 泛型版本的 Map Reduce Filter
func Demo() {
	list01 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result01 := ComputableMap(list01, func(a int) int {
		return a * a
	})
	fmt.Println(result01)

	result02 := ComputableReduce(list01, func(a int) int {
		if a > 2 {
			return 0
		}
		return a
	})
	fmt.Println(result02)

	result03 := ComparableFilter(list01, func(a int) bool {
		return a > 3
	})
	fmt.Println(result03)

	list02 := []string{"123", "23", "3"}
	result04 := ComparableFilter(list02, func(a string) bool {
		return len(a) > 1
	})
	fmt.Println(result04)
}

type ComputableTypes interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func ComputableMap[T ComputableTypes](list []T, fn func(T) T) []T {
	ret := make([]T, 0, len(list))
	for _, item := range list {
		ret = append(ret, fn(item))
	}
	return ret
}

func ComputableReduce[T ComputableTypes](list []T, fn func(T) T) T {
	var sum T
	for _, item := range list {
		sum += fn(item)
	}
	return sum
}

type ComparableTypes interface {
	int | string | bool
}

func ComparableFilter[T ComparableTypes](list []T, fn func(T) bool) []T {
	ret := make([]T, 0)
	for _, item := range list {
		if fn(item) {
			ret = append(ret, item)
		}
	}
	return ret
}
