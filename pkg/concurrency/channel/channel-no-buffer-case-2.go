package main

import (
	"fmt"
	"time"
)

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func(){
		println("worker start to work...")
		// 调用worker
		f()
		// 接收动作：等待接收，阻塞，直至有发送动作
		c <- signal(struct{}{})
	}()

	return c
}

func main() {
	println("start a worker...")
	// spawn函数返回的channel被用于承载新goroutine退出的通知信号
	// 该信号专用于通知main goroutine
	c := spawn(worker)
	// 发送动作
	<-c
	fmt.Println("worker work done!")
}