package main

import (
	"fmt"
	"time"
)

func add(p *int32) {
	*p++
}

func main() {
	c := int32(0)
	for i := 0; i < 1000; i++ {
		go add(&c)
	}

	time.Sleep(time.Second)
	fmt.Println(c)
}
