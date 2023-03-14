package main

import "fmt"

func b() {
	i := 0
	// 1.18 之前
	// 如果函数参数为inteface{}
	// 函数的实参很可能逃逸
	// 因为interface{}类型的函数往往会使用反射
	fmt.Println(i)
}
func main() {
	b()
}
