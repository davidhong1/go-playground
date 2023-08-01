package main

// 非线程安全
type IntSet struct {
	data map[int]bool
}

// IntSet 实现了 Add()、Delete()、Contains() 三个基础操作
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Len() int {
	return len(set.data)
}
