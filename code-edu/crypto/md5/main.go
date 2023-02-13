package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Printf("%x\n", md5.Sum([]byte("测试数据")))
}
