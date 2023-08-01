package main

import (
	"fmt"
)

func main() {
	is1 := NewIntSet()
	is1.Add(1)
	is1.Add(2)
	is1.Delete(2)
	if is1.Contains(2) {
		fmt.Println("Contain 2")
	} else {
		fmt.Println("Don't contain 2")
	}

	is2 := NewUndoableIntSet()
	is2.Add(1)
	is2.Add(2)
	is2.Add(3)
	fmt.Println(is2.Len()) // 3
	is2.Undo()
	fmt.Println(is2.Len()) // 2

	is3 := NewIOCDIPIntSet()
	is3.Add(1)
	is3.Add(2)
	is3.Add(3)
	fmt.Println(is3.Len()) // 3
	is3.Undo()
	fmt.Println(is3.Len()) // 2
}
