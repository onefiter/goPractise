package demo

import (
	"fmt"
	"sync"
	"testing"
)

// 同步锁

var (
	count int
	wg    sync.WaitGroup
	mu    sync.Mutex
)

// 当一个goroutine需要访问共享资源时，它可以首先调用Lock()方法来获取锁，
// 只有成功获取锁的goroutine才能进入临界区访问共享资源。
// 其他goroutine在没有获取到锁的情况下，会被阻塞在Lock()方法处，直到锁被释放。
func TestSynchronized(t *testing.T) {
	wg.Add(2)
	go increment()
	go increment()

	wg.Wait()
	fmt.Println(" Final count:", count)
}

func increment() {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}
