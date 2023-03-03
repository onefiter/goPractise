package main

var c = make(chan int)

var a string

func f() {
	a = "hello, world"
	// 发送动作
	<-c
}

func main() {
	go f()
	// 接收动作
	c <- 5
	println(a)
}