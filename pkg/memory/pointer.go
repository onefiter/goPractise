package main

import "fmt"

// 指针逃逸
// 函数返回了指针
func a() *int {
	v := 0
	fmt.Printf("%p\n", &v)
	return &v
}

func main() {
	i := a()

	fmt.Printf("%p\n", i)
}
