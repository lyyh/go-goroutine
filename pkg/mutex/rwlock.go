package mutex

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(1 * time.Millisecond) // 假设读操作耗时10毫秒
	fmt.Println("write", x)
	rwlock.Unlock() // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	fmt.Println("read", x)
	rwlock.RUnlock() // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func Exec() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println("write for traverel")
		wg.Add(1)
		go write()
	}
	// 一直进读锁导致写锁进不了
	for i := 0; i < 100; i++ {
		fmt.Println("read for traverel")
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
