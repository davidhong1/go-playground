package main

import (
	"errors"
	"fmt"
)

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		fmt.Println("undo, index:", index)
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

// 不是由控制逻辑 Undo 来依赖业务逻辑 IntSet，
// 而是由业务逻辑 IntSet 依赖 Undo。
// 这样一来，我们 Undo 代码就可以复用了。
type IOCDIPIntSet struct {
	IntSet
	undo Undo
}

func NewIOCDIPIntSet() IOCDIPIntSet {
	return IOCDIPIntSet{
		IntSet: NewIntSet(),
		undo:   nil,
	}
}
func (set *IOCDIPIntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IOCDIPIntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IOCDIPIntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
