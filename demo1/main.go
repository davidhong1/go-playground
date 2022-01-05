package main

import (
	"demo1/something"
	"fmt"
)

func main() {
	// fmt.Println("hello world")

	bes := "test"
	afs := something.ReverseRunes(bes)
	fmt.Printf("%s reverse to %s\n", bes, afs)
}
