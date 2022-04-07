package main

import (
	"fmt"
	"unsafe"
)

func main() {

}

func main1() {
	bytes := []byte{104, 101, 108, 108, 111}

	p := unsafe.Pointer(&bytes) // 强制转换成 unsafe.Pointer，编译器不会报错

	// 因为 string 底层也是 []byte，两者具有相同的 layout
	str := *(*string)(p) // 然后强制转换成 string 类型的指针，再将这个指针的值当做 string 类型取出来

	fmt.Println(str) // 输出 hello
}
