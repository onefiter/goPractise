package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s := "斗破苍穹"
	s2 := "斗破苍穹2abc"
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(s2))

	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	// uicode 三个字节表示一个中文，所以共12个字节
	// 所有的字符均使用Unicode字符集
	// Unicode
	// 1. 一种统一的字符集
	// 2. 囊括了159种文字的144679个字符
	fmt.Println(sh.Len)
}
