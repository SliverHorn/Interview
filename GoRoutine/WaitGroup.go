package GoRoutine

import (
	"fmt"
	"sync"
)
/*
WaitGroup
- 理解WaitGroup的实现,核心是==CAS==的使用
- Add与Done应该放在哪里?
	- Add放在goroutine外
	- Done放在goroutine中
	- 逻辑复杂是建议使用defter保证调用
 */

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