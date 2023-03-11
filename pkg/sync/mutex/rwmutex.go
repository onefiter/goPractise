package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.RWMutex
	level  int
	salary int
}

func (p *Person) printPerson() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println(p.salary)
	fmt.Println(p.level)
}

func main() {
	p := Person{level: 1, salary: 10000}

	go p.printPerson()
	go p.printPerson()
	go p.printPerson()

	time.Sleep(time.Second)

}
