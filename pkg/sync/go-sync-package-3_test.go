package main

import (
	"sync"
	"testing"
)
// sync包提供了两种用于临界区同步的原语
// 1. 互斥锁
// 2. 读写锁
// 互斥锁是临界区同步原语的首选，它常被用来对结构体对象的内部状态、缓存等进行保护，是使用最为广泛的临界区同步原语。
// 读写锁有其存在的道理和适用场景。

// 模拟临界区要保护的数据
var cs1 = 0
var mu1 sync.Mutex

// 模拟临界区要保护的数据
var cs2 = 0
var mu2 sync.RWMutex



func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			mu1.Lock()
			_ = cs1
			mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			mu2.RLock()
			_ = cs2
			mu2.RUnlock()
		}
	})	
}

func BenchmarkWriteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			mu2.Lock()
			_ = cs2
			mu2.Unlock()
		}
	})	
}