package main

import (
	"fmt"
	"unsafe"
)

// 空结构体size大小是0

// runtime/malloc.go 大概777行
// base address for all 0-byte allocations
// var zerobase uintptr

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
