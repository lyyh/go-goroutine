package mutex

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x = 0

// 互斥锁保证一个时间只有一个goroutine进入临界区
// 多个goroutine在等待同一个锁时，唤醒的策略是随机的
func add() {
	for i := 0; i < 100; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func init() {
	x = 1
	wg.Add(2)

	go add()
	go add()
	wg.Wait()
	fmt.Println("ok", x)
}
