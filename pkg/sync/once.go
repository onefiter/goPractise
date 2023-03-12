package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.Mutex
	level  int
	salary int
}

func (p *Person) promote() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.salary += 1000
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)

}

func main() {
	p := Person{level: 1, salary: 10000}

	// p.promote只执行一次
	once := sync.Once{}
	go once.Do(p.promote)
	go once.Do(p.promote)
	go once.Do(p.promote)

	time.Sleep(time.Second)

}
