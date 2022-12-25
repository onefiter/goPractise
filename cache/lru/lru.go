package lru

import (
	"container/list"

	"github.com/goPractise/cache"
)

// lru 是一个LRU cache。它不是并发安全的。
type lru struct {
	// 缓存最大的容量，单位字节
	// groupcache 使用的是最大存放 entry 个数
	maxBytes int
	// 当一个entry 从缓存中移除是调用该回调函数，默认是nil
	// groupcache 中的key是任意的可比较类型；value是interface{}
	onEvicted func(key string, value interface{})

	// 已使用的字节数，只包括值，key不算
	usedBytes int

	ll    *list.List
	cache map[string]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CaclLen(e.value)
}

// 创建一个新的Cache， 如若maxBytes是0，表示没有容量限制

func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &lru{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}

func (l *lru) Set(key string, value interface{}) {
	if e, ok := l.cache[key]; ok {
		l.ll.MoveToBack(e)
		en := e.Value.(*entry)
		l.usedBytes = l.usedBytes - cache.CaclLen(en.value) + cache.CaclLen(value)
		en.value = value
		return
	}

	en := &entry{key, value}
	e := l.ll.PushBack(en)
	l.cache[key] = e

	l.usedBytes += en.Len()

	if l.maxBytes > 0 && l.usedBytes > l.maxBytes {

		l.DelOldest()

	}

}

func (l *lru) Get(key string) interface{} {
	if e, ok := l.cache[key]; ok {
		l.ll.MoveToBack(e)
		return e.Value.(*entry).value
	}

	return nil
}

func (l *lru) Del(key string) {
	if e, ok := l.cache[key]; ok {
		l.removeElement(e)
	}
}

func (l *lru) DelOldest() {
	l.removeElement(l.ll.Front())
}

func (l *lru) Len() int {
	return l.ll.Len()
}

func (l *lru) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	l.ll.Remove(e)
	en := e.Value.(*entry)
	l.usedBytes -= en.Len()

	delete(l.cache, en.key)

	if l.onEvicted != nil {
		l.onEvicted(en.key, en.value)
	}
}
