package main

import (
	"fmt"
	"time"
)

// 通过共享内存来传递数据
func watch(p *int) {
	for {
		if *p == 1 {
			fmt.Println("hello")
			break
		}
	}

}

func main() {
	i := 0

	go watch(&i)

	time.Sleep(time.Second)

	i = 1

	time.Sleep(time.Second)
}
