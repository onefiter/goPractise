package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func add(p *int32) {
	// *p++
	// 实现的是CPU级别的锁,硬件锁

	atomic.AddInt32(p, 1)
}

func main() {
	c := int32(0)
	for i := 0; i < 1000; i++ {
		go add(&c)
	}

	atomic.CompareAndSwapInt32()

	time.Sleep(time.Second)
	fmt.Println(c)
}
