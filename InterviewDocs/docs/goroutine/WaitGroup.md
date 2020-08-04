# WaitGroup
- 理解WaitGroup的实现,核心是 `CAS` 的使用
- Add与Done应该放在哪里?
	- Add放在goroutine外
	- Done放在goroutine中
	- 逻辑复杂是建议使用defter保证调用
- WaitGroup适合什么样的场景
    - 并发的goroutine执行的逻辑相同时, 否则代码并不简介,可以采用其他方式
	
```go
package main

import (
	"fmt"
	"sync"
)

// WaitGroup是一个结构体,不能copy是因为函数是值拷贝
func ErrorWaitGroup1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg sync.WaitGroup) {
			fmt.Println(i)
			defer wg.Done()
		}(i, wg)
	}
}


// WaitGroup的Add要在goroutine前执行
func ErrorWaitGroup2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
}

// WaitGroup的Add 很大会有影响吗?
func ErrorWaitGroup3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(100)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
			wg.Add(-100)
		}(i)
	}
	fmt.Println("Finished")
}

func main() {
	ErrorWaitGroup1()
	ErrorWaitGroup2()
	ErrorWaitGroup3()
}
```