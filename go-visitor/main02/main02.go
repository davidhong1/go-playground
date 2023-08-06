package main02

import "fmt"

// Visitor 修饰器模式
func Demo02() {
	fmt.Println("------Demo02------")
	info := Info{}
	var v Visitor = &info
	v = NewDecoratedVisitor(v, WithNameVisitor, WithOtherThingsVisitor)

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
	fmt.Println("------Demo02------")
}

func WithNameVisitor(fn VisitorFunc) VisitorFunc {
	return func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call funciton")
		err = fn(info, err) // a. 首先执行传入的 Visitor
		if err == nil {
			// 没有报错，则输出主要逻辑
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	}
}

func WithOtherThingsVisitor(fn VisitorFunc) VisitorFunc {
	return func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	}
}

type VisitorDecorator func(VisitorFunc) VisitorFunc

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorDecorator
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorDecorator) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor.
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	decoratorLen := len(v.decorators)
	for i := range v.decorators {
		d := v.decorators[decoratorLen-i-1]
		fn = d(fn)
	}
	return fn(v.visitor.(*Info), nil)
}
