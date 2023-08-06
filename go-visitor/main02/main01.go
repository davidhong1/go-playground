package main02

import "fmt"

// 解耦了数据和程序
// 使用了修饰器模式
// 还做出了 pipeline 的模式
func Demo01() {
	fmt.Println("------Demo01------")
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
	fmt.Println("------Demo01------")
}

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

// Name Visitor
// 这个 Visitor 主要是用来访问 Info 结构中的 Name 和 NameSpace 成员
type NameVisitor struct {
	// 用于嵌套 Visitor
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit( // c. 然后执行嵌套的 Visitor 逻辑
		func(info *Info, err error) error { // 这个是 NameVisitor 定义的逻辑 b. 执行本 Visitor 定义的逻辑
			fmt.Println("NameVisitor() before call funciton")
			err = fn(info, err) // a. 首先执行传入的 Visitor
			if err == nil {
				// 没有报错，则输出主要逻辑
				fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
			}
			fmt.Println("NameVisitor() after call function")
			return err
		},
	)
}

// Other Visitor
// 这个 Visitor 主要用于访问 Info 结构中的 OtherThings 成员
type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}
