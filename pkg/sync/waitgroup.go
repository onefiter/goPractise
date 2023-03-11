package main

import (
	"fmt"
	"sync"
)

type Person struct {
	mu     sync.RWMutex
	level  int
	salary int
}

func (p *Person) printPerson(w *sync.WaitGroup) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println(p.salary)
	fmt.Println(p.level)
	w.Done()
}

func main() {
	p := Person{level: 1, salary: 10000}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go p.printPerson(&wg)
	go p.printPerson(&wg)
	go p.printPerson(&wg)
	wg.Wait()

}
