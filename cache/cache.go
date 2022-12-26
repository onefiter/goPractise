package cache

import (
	"log"
	"sync"
)

// Cache缓存接口
type Cache interface {
	Set(key string, value interface{}) // 设置/添加一个缓存，如果key存在,则用新值覆盖
	Get(key string) interface{}        // 通过key获取一个缓存值
	Del(key string)                    // 根据key删除一个缓存值
	DelOldest()                        // 删除最“无用”的一个缓存值
	Len() int                          // 获取缓存已存在的记录数
}

type Value interface {
	Len() int
}

// DefaultMaxBytes 默认允许占用的最大内存
const DefaultMaxBytes = 1 << 29

// safeCache 并发安全缓存
type safeCache struct {
	m     sync.RWMutex
	cache Cache

	nhit, nget int
}

func newSafeCache(cache Cache) *safeCache {
	return &safeCache{
		cache: cache,
	}
}

func (sc *safeCache) set(key string, value interface{}) {
	sc.m.Lock()
	defer sc.m.Unlock()
	sc.cache.Set(key, value)
}

func (sc *safeCache) get(key string) interface{} {
	sc.m.RLock()
	defer sc.m.RUnlock()
	sc.nget++
	if sc.cache == nil {
		return nil
	}
	v := sc.cache.Get(key)
	if v != nil {
		log.Println("[TourCache] hit")
		sc.nhit++
	}

	return v
}

func (sc *safeCache) stat() *Stat {
	sc.m.RLock()
	defer sc.m.RUnlock()
	return &Stat{
		NHit: sc.nhit,
		NGet: sc.nget,
	}
}

type Stat struct {
	NHit, NGet int
}
