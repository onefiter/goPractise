package main

import (
	"fmt"
	"time"
)

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	fmt.Println("调用了f3")
}

func main() {
	go f1()

	time.Sleep(10 * time.Second)
}
