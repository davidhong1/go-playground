package main02

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 0, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func Demo() {
	// 1. 统计有多少员工大于 40 岁
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("Old People: %d\n", old)
	// 2. 统计有多少员工的薪水大于 6000
	highPay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary > 6000
	})
	fmt.Printf("High Salary People: %d\n", highPay)
	// 3. 列出有没有休假的员工
	noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People No Vacation: %v\n", noVacation)
	// 4. 统计所有员工的薪资总和
	totalPay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", totalPay)
	// 5. 统计 30 岁以下员工的薪资总和
	youngerPay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("Younger Total Salary: %d\n", youngerPay)
}

// EmployeeCountIf 和 EmployeeSumIf 分别用于统计满足某个条件的个数或总数。
// 它们都是 Filter + Reduce 的语义。
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i := range list {
		sum += fn(&list[i])
	}
	return sum
}

// EmployeeFilterIn 就是按某种条件过滤，就是 Filter 的语义。
func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}
