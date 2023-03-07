package main

import (
	"fmt"
	"unsafe"
)

// 空结构体size大小是0

type S struct {
}

type K struct{}

func main() {

	s := S{}
	k := K{}

	fmt.Printf("size of int: %d\n", unsafe.Sizeof(int(0)))

	fmt.Printf("size of empty struct: %d\n", unsafe.Sizeof(s))

	// zerobase 在runtime.go中 runtime/malloc.go
	fmt.Printf("the address of empty struct: %p\n", &s)
	fmt.Printf("the address of empty struct: %p\n", &k)

}
