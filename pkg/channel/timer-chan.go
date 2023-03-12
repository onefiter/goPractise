package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second)

	<-t.C

	fmt.Println("hello")
}
