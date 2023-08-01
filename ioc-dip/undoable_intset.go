package main

import (
	"errors"
	"fmt"
)

// Poor style(不够好的风格)
// 在这里业务逻辑中写了控制逻辑，
// 需要业务逻辑和控制分离，
// 让业务逻辑依赖控制逻辑。
type UndoableIntSet struct {
	IntSet    // Embedding(delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		fmt.Println("undo, index:", index)
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}
