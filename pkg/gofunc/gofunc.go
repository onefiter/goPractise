package main

import "time"

func main() {
	go func() {
		println("hello world in goroutine")
	}()

	time.Sleep(time.Second * 5)
}
