package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int, ch chan struct{}) {
	fmt.Println(i)
	time.Sleep(time.Second)
	<-ch
}

// 这里会牺牲一些性能，但两权相害，取其轻
func main() {
	c := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt32; i++ {
		c <- struct{}{}
		go do(i, c)
	}

	time.Sleep(time.Hour)
}
