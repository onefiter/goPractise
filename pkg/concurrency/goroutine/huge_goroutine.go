package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int) {
	fmt.Println(i)
	time.Sleep(time.Second)
}

// panic: too many concurrent operations on a single file or socket (max 1048575)
func main() {
	for i := 0; i < math.MaxInt32; i++ {
		go do(i)
	}

	time.Sleep(time.Hour)
}
