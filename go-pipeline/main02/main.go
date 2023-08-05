package main02

import (
	"fmt"
	"math"
	"sync"
)

// Fan In/Out
// 动用 Go 语言的 Go Routine 和 Channel 还有一个好处，就是可以写出 1 对多，或多对 1 的 pipeline，
// 也就是 Fan In / Fan Out
func Demo() {
	// 首先，我们制造了从 1 到 10000 的数组;
	nums := makeRange(1, 10000)
	// 然后，把这堆数组全部 echo 到一个 channel 里 in;
	in := echo(nums)
	// 此时，生成 5 个 channel，接着都调用 sum(prime(in))，
	// 于是，每个 sum 的 go routine 都会开始计算和
	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		// 5 个并发
		chans[i] = sum(prime(in))
	}
	// 最后，再把所有结果求和拼起来，得到最终结果
	// merge 起到等待 chans[i] = sum(prime(in)) 执行完的动作
	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if isPrime(n) {
				// fmt.Println("prime:", n)
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 是否是质数
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
