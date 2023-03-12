package main

import "fmt"

// 非阻塞使用channel
func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int)

	select {
	case <-c1:
		fmt.Println("c1")
	case c2 <- 1:
		fmt.Println("c2")
	default:
		fmt.Println("none")
	}
}
