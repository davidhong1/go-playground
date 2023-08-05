package main01

import "fmt"

func Demo() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)
	}

	// 上面的代码类似于我们执行了 Unix/Linux 命令
	// echo $nums | sq | sum
	// 如果你不想有那么多的函数嵌套，就可以使用一个代理函数来完成
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func echo(nums []int) <-chan int {
	out := make(chan int) // 通道缓存区只有 1 个
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // close 后，只能从 close 中获取数据
	}()
	return out
}

// 平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		// 同步等待最后的 sum
		out <- sum
		close(out)
	}()
	return out
}
