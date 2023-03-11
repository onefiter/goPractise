package main

import (
	"fmt"
	"time"
)

// 通过共享内存来传递数据
func watch(c chan int) {

	if <-c == 1 {
		fmt.Println("hello")
	}

}

func main() {
	// 不要通过共享内存的方式进行通信
	// 而是应该通过通信的方式共享内存
	c := make(chan int)

	go watch(c)

	time.Sleep(time.Second)

	c <- 1

	time.Sleep(time.Second)
}
