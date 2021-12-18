package concurrency

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func task(index int) {
	defer wg.Done()
	fmt.Printf("执行了一个goroutine %d\n", index)
}
func WaitRun() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i) // 并发执行的，系统随机调度不能保证顺序
	}
	wg.Wait()
}
