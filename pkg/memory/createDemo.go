package main

import "fmt"

type Demo struct {
	name string
}

func createDemo(name string) *Demo {
	//局部变量d 逃逸到堆
	d := new(Demo)
	d.name = name
	return d

}

func main() {
	demo := createDemo("demo")
	fmt.Println(demo)
}
