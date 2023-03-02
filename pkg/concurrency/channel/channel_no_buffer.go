package main

// 无缓冲channel兼具通信和同步特性，在并发程序中应用颇为广泛。
// 可以通过不带有capacity参数的内置make函数创建一个可用的无缓冲channel
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
	println("aaa")
}

func main() {
	go f()

	c <- 5

	println(a)
}
