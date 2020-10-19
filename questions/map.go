package main


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
可以这样理解，假设两个goroutine：g1和g2，
当g1 load到同学A的值为x，即将要将x+1 store但还没有store的时候，
g2 load也load到同学A的值为x,并把x+1store成功，
此时g1再去将x+1 store的时候就会导致同学A的值失真
sync.Map 本身是并发安全的，但是你对它的操作不是并发安全的。
*/

var (
	resultNum   int
	tmp         interface{}
	wg          sync.WaitGroup
	rankingList sync.Map
	lock        sync.Mutex
)

// randomAndAdd 是进行数字1-3的随机摇号，并根据摇号结果，将对应编号的同学得票次数+1的方法
func randomAndAdd() {
	// 每次生成新的随机数
	rand.Seed(time.Now().UnixNano())

	resultNum = rand.Intn(3) + 1

	// lock.Lock()
	switch resultNum {
	case 1:
		tmp, _ = rankingList.Load("同学A")
		rankingList.Store("同学A", tmp.(int)+1)
	case 2:
		tmp, _ = rankingList.Load("同学B")
		rankingList.Store("同学B", tmp.(int)+1)
	case 3:
		tmp, _ = rankingList.Load("同学C")
		rankingList.Store("同学C", tmp.(int)+1)
	}
	// lock.Unlock()
	wg.Done()

}

func main() {
	// 生成key为0的初始表，包含3个同学
	rankingList.Store("同学A", 0)
	rankingList.Store("同学B", 0)
	rankingList.Store("同学C", 0)

	// 重复启动1万个goroutine
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go randomAndAdd()
	}
	wg.Wait()

	// sum是统计所有同学总票数的变量
	sum := 0

	// 循环输出每个同学的得票数，并将每个同学的得票数累加到sum
	rankingList.Range(func(k, v interface{}) bool {
		sum += v.(int)
		fmt.Println(k, v)
		return true
	})

	// 输出所有同学总票数
	fmt.Printf("所有同学总票数=%v\n", sum)
}