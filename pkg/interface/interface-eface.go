package main

import "fmt"

// eface
type T struct {
	n int
	s string
}

// iface
type T2 struct {
	n int
	s string
}

func (T2) M1() {}
func (T2) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T{
		n: 17,
		s: "hello, interface",
	}
	// Go 运行时使用eface结构表示ei
	var ei interface{} = t

	fmt.Println(ei)

	// 此时t2是NonEmptyInterface变量类型
	var t2 = T2{
		n: 18,
		s: "hello, interface2",
	}

	var i NonEmptyInterface = t2
	fmt.Printf("the type of value i: %T\n", i)
	fmt.Println(i)
}
